#!/usr/bin/env python
from typing import List
import json

from sagas.modules.deles import *

def keymaps(rel):
    return [{"fieldName":k.getFieldName(),
             "relFieldName": k.getRelFieldName(),
             }
            for k in rel.getKeyMaps()]

class MetaGenerator(object):
    def abi(self, entities:List[str]):
        """
        $ python meta_generator.py abi [InventoryItemAndDetail]
        $ python meta_generator.py abi [Person]
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
            relations=[{"name":rel.getCombinedName(), "type": rel.getType(),
                        "relEntityName": rel.getRelEntityName(),
                        "fkName": rel.getFkName(),
                        "keymaps": keymaps(rel)
                        }
                       for rel in model.getRelations()]
            abi={
                "name": model.getEntityName(),
                "fields": fields,
                "relations": relations
            }

            out_file=open(ent.lower()+".json", 'w')
            out_file.write(json.dumps(abi, indent=2, ensure_ascii=False))
            out_file.close()

if __name__ == '__main__':
    import fire
    fire.Fire(MetaGenerator)
