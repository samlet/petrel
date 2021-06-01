#!/usr/bin/env python
from typing import List

from sagas.modules.deles import *
import json

def filter_params(model, prec):
    params=[{"name": param.getName(),
             "type": param.getType(),
             "required": param.isOptional(),
             "overrideOptional": param.isOverrideOptional(),
             "entity": param.getEntityName(),
             "mode": param.getMode(),
             "internal": param.isInternal()
             } for param in model.getModelParamList()
            if prec(param.getMode())]
    params= [p for p in params if not p['internal']]
    return list({v['name']:v for v in params}.values())

class ServiceMeta(object):
    def entity_services(self, ent):
        """
        $ python service_meta.py entity_services Person
        :param ent:
        :return:
        """
        services=oc.all_service_names()
        result=[]
        for serv_name in services:
            model_serv = oc.service_model(serv_name)
            def_ent = model_serv.getDefaultEntityName()
            if def_ent==ent:
                result.append(serv_name)

        return result

    def abi(self, service):
        model=MetaService(service).model
        service_def={
            "name": service,
            "description": model.getDescription(),
            "paramsInput": filter_params(model, lambda m: m=="IN" or m=="INOUT"),
            "paramsOutput": filter_params(model, lambda m: m=="OUT" or m=="INOUT")
        }

        return service_def

    def entity_abi(self, ent, extra_services:List[str]=None):
        """
        $ python service_meta.py entity_abi Person
        $ python service_meta.py entity_abi Example [createExampleStatus]
        :param ent:
        :return:
        """
        srvs=self.entity_services(ent)
        service_defs=[]
        for srv in srvs:
            service_defs.append(self.abi(srv))
        if extra_services is not None:
            service_defs.extend([self.abi(srv) for srv in extra_services])

        abi={"entity": ent,
             "package": ent.lower(),
             "ops": service_defs}
        out_file=open("assets/"+ent.lower()+"_ops.json", 'w')
        out_file.write(json.dumps(abi, indent=2, ensure_ascii=False))
        out_file.close()

if __name__ == '__main__':
    import fire
    fire.Fire(ServiceMeta)

