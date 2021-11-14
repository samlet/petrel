#!/usr/bin/env python
from typing import List

import io_utils
from sagas.modules.deles import *
import json
import os

templates = {
    "digest": """
// {{srv.description}}
service {{srv.name}}:
    {% for para in srv.paramsInput %}
    field: {{para.name}} ({{para.type}})
    {%- endfor %}
}
""",
    "class": """
// {{srv.description}}
@Data
public class {{srv.name}} {
    {%- for para in srv.paramsInput %}
        private {{para.type}} {{para.name}}; 
    {%- endfor %}
}
""",
# task agent
    "agent": """
package com.bluecc.prefabs.services;

import com.bluecc.prefabs.fixtures.ServiceBase;
import com.bluecc.prefabs.fixtures.ServiceCaller.*;
import com.bluecc.prefabs.kafka.Receiver;
import com.google.common.collect.Maps;
import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.Singular;
import java.util.Map;

// {{srv.description}}
public class {{srv.className}}Meta {
    
    @Data
    @Builder
    public static class {{srv.className}}Params{
     {%- for para in srv.paramsInput %}
        private {{para.type}} {{para.name}}; //{% if para.required %} required {% endif %}
    {%- endfor %}
    }

    @Data
    @Builder
    @EqualsAndHashCode(callSuper = true)
    static class {{srv.className}} extends ServiceBase{
        {{srv.className}}Params params;

        public String send(CallContext ctx, Receiver.ReceiveCallback callback) {
            ServiceEnvelope<{{srv.className}}Params> msg = ServiceEnvelope.<{{srv.className}}Params>builder()
                    .serviceName("{{srv.name}}")
                    .serviceInParams(params)
                    .build();

            return sendImpl(ctx, callback, msg);
        }
    }    
}    
""",
}

from dotenv import load_dotenv
load_dotenv()  # take environment variables from .env.

def filter_params(model, prec):
    params = [{"name": param.getName(),
               "type": param.getType(),
               "required": not param.isOptional(),
               "overrideOptional": param.isOverrideOptional(),
               "entity": param.getEntityName(),
               "mode": param.getMode(),
               "internal": param.isInternal()
               } for param in model.getModelParamList()
              if prec(param.getMode())]
    params = [p for p in params if not p['internal']]
    return list({v['name']: v for v in params}.values())


class ServiceMeta(object):
    def entity_services(self, ent):
        """
        $ python service_meta.py entity_services Person
        :param ent:
        :return:
        """
        services = oc.all_service_names()
        result = []
        for serv_name in services:
            model_serv = oc.service_model(serv_name)
            def_ent = model_serv.getDefaultEntityName()
            if def_ent == ent:
                result.append(serv_name)

        return result

    def abi(self, service):
        """
        $ python service_meta.py abi createPerson
        :param service:
        :return:
        """
        # from sagas.util.str_converters import to_camel_case
        model = MetaService(service).model
        service_def = {
            "name": service,
            # "className": to_camel_case(service, True),
            "className": service[0].capitalize()+service[1:],
            "description": model.getDescription(),
            "paramsInput": filter_params(model, lambda m: m == "IN" or m == "INOUT"),
            "paramsOutput": filter_params(model, lambda m: m == "OUT" or m == "INOUT")
        }

        return service_def

    def write_abi(self, service: str, tpl=None, display=True):
        """
        $ python service_meta.py write_abi createPerson
        $ python service_meta.py write_abi createPerson digest

        * 如果.env文件中定义了SERVICE_META变量, 则会输出到该目录下
        $ python service_meta.py write_abi updatePerson agent

        :param service:
        :return:
        """
        from template_helper import render
        meta = self.abi(service)
        if display:
            cnt = json.dumps(meta, ensure_ascii=False, indent=2)
            print(cnt)

        # bind templates
        if tpl is not None:
            cnt = render(templates[tpl], srv=meta)
            if display:
                print(cnt)

            srv_meta_dir=os.getenv("SERVICE_META")
            if srv_meta_dir is not None:
                file=f"{srv_meta_dir}/{meta['className']}Meta.java"
                io_utils.write_to_file(file, cnt)
                print('write to', file)

    def write_all_abi(self, cfg_file:str):
        """
        $ python -m service_meta write_all_abi srvs_party.txt
        $ python -m service_meta write_all_abi srvs_example.txt

        :param cfg_file:
        :return:
        """
        names=io_utils.read_file(cfg_file).split(",")
        names=[n.strip() for n in names]
        print(names)
        for srv in names:
            self.write_abi(srv, 'agent', display=False)

    def entity_abi(self, ent, extra_services: List[str] = None):
        """
        $ python service_meta.py entity_abi Person
        $ python service_meta.py entity_abi Example [createExampleStatus]
        :param ent:
        :return:
        """
        srvs = self.entity_services(ent)
        service_defs = []
        for srv in srvs:
            service_defs.append(self.abi(srv))
        if extra_services is not None:
            service_defs.extend([self.abi(srv) for srv in extra_services])

        abi = {"entity": ent,
               "package": ent.lower(),
               "ops": service_defs}
        out_file = open("assets/" + ent.lower() + "_ops.json", 'w')
        out_file.write(json.dumps(abi, indent=2, ensure_ascii=False))
        out_file.close()


if __name__ == '__main__':
    import fire

    fire.Fire(ServiceMeta)
