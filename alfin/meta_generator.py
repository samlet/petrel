#!/usr/bin/env python
from typing import List
import json
from os import listdir
from os.path import isfile, join

import xml.etree.ElementTree as ET

from sagas.modules.deles import *
from utils import create_dir, write_to_file

def keymaps(rel):
    return [{"fieldName":k.getFieldName(),
             "relFieldName": k.getRelFieldName(),
             }
            for k in rel.getKeyMaps()]


def collect_types(seed_path, f):
    xml_path=join(seed_path, f)
    tree = ET.parse(xml_path)
    root = tree.getroot()
    ent_types=set([c.tag for c in root])
    return ent_types

def collect_data(seed_path, f):
    rs=[]
    xml_path=join(seed_path, f)
    tree = ET.parse(xml_path)
    root = tree.getroot()
    for c in root:
        if c.tag not in [
            'create','create-replace',
            'create-update',
            'delete']:
            rs.append({c.tag:c.attrib})
    return rs

class MetaGenerator(object):
    def __init__(self, asset_root=""):
        self.asset_dir="assets"
        self.asset_root=asset_root

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
                        "keymaps": keymaps(rel),
                        "autoRelation": rel.isAutoRelation(),
                        }
                       for rel in model.getRelations()]
            abi={
                "name": model.getEntityName(),
                "fields": fields,
                "relations": relations,
                "pksSize": model.getPksSize(),
                "pks": [f for f in model.getPkFieldNames()]
            }

            target_dir=self.asset_dir
            if self.asset_root!="":
                target_dir=f"{target_dir}/{self.asset_root}/"
                create_dir(target_dir)
            out_file=open(target_dir+"/"+ent.lower()+".json", 'w')
            out_file.write(json.dumps(abi, indent=2, ensure_ascii=False))
            out_file.close()

    def list_ents(self, asset_path='data_example', write=False):
        """
        $ python meta_generator.py list_ents data_example

        :param asset_path:
        :param write:
        :return:
        """
        seed_path=f"./{self.asset_dir}/{asset_path}"
        sets = self.collect_entity_types(seed_path)

        if write:
            fd=open('{self.asset_dir}/data_example_ents.json', 'w')
            fd.write(json.dumps(sets, indent=2))
            fd.close()
        else:
            print(f'total entities {len(sets)}:', sets)

    def collect_entity_types(self, seed_path):
        onlyfiles = [f for f in listdir(seed_path) if isfile(join(seed_path, f))]
        all_types = set([])
        for f in onlyfiles:
            if f.endswith('.xml'):
                types = collect_types(seed_path, f)
                print(f"{f} contains {len(types)}")
                all_types.update(types)
        sets = [t for t in all_types if t not in [
            'create', 'create-replace',
            'create-update',
            'delete']]
        return sets

    def gen_seed_abi(self, asset_path='data_example'):
        """
        $ python meta_generator.py gen_seed_abi data_example
        :param asset_path:
        :return:
        """
        seed_path=f"{self.asset_dir}/{asset_path}"
        sets = self.collect_entity_types(seed_path)
        self.abi(sets)

    def convert_data(self, asset_path='data_example'):
        """
        $ python meta_generator.py convert_data data_example
        $ python meta_generator.py convert_data demo
        :param asset_path:
        :return:
        """
        import os

        seed_path=f"./{self.asset_dir}/{asset_path}"
        onlyfiles = [f for f in listdir(seed_path) if isfile(join(seed_path, f))]
        for f in onlyfiles:
            if f.endswith('.xml'):
                rs=collect_data(seed_path, f)
                target_path=join(seed_path, os.path.splitext(f)[0]+".json")
                fd=open(target_path, 'w')
                fd.write(json.dumps(rs, indent=2, ensure_ascii=False))
                fd.close()

    def gen_package(self, pkg, asset_root):
        """
        $ python meta_generator.py gen_package "com.bluecc.workload" workload
        :param pkg:
        :param asset_root:
        :return:
        """
        self.asset_root=asset_root

        reader=oc.delegator.getModelReader()
        filters=oc.jset(pkg)
        ent_map=reader.getEntitiesByPackage(filters, oc.jset())
        all_ents=[]
        for k,v in ent_map.items():
            all_ents.extend(v)

        self.abi(all_ents)

        # write package meta
        pkg_meta={"name":self.asset_root,
                  "package": pkg,
                  "entities": {ent:ent.lower()+".json" for ent in all_ents}
                  }
        write_to_file(f"{self.asset_dir}/{self.asset_root}/meta.json",
                      json.dumps(pkg_meta, indent=2))

        print("ok.")

if __name__ == '__main__':
    import fire
    fire.Fire(MetaGenerator)
