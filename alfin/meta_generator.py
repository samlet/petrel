#!/usr/bin/env python
from typing import List
import json

from sagas.modules.deles import *

class MetaGenerator(object):
    def abi(self, entities:List[str]):
        """
        $ python meta_generator.py fields InventoryItemAndDetail
        :param ent:
        :return:
        """
        for ent in entities:
            model=oc.delegator.getModelEntity(ent)
            fields=[{"name":f.getName(), "type":f.getType(), "col":f.getColName(),
                     "pk":f.getIsPk(), "notNull":f.getIsNotNull(), "encrypt":f.getEncrypt(),
                     "autoCreatedInternal": f.getIsAutoCreatedInternal(),
                     "validators":[v for v in f.getValidators()]
                     }
                    for f in model.getFieldsIterator()]
            abi={
                "name": model.getEntityName(),
                "fields": fields
            }

            out_file=open(ent.lower()+".json", 'w')
            out_file.write(json.dumps(abi, indent=2, ensure_ascii=False))
            out_file.close()

if __name__ == '__main__':
    import fire
    fire.Fire(MetaGenerator)
