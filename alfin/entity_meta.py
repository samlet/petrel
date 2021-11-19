import json
from os.path import join

import io_utils

templates = {
    "simple": '''
{{ent.name}}
fields:
{%- for fld in ent.fields %}
    - {{fld.name}} ({{fld.type}})  
{%- endfor %}

relations:
{%- for rel in ent.relations %}
    - {{rel.name}} ({{rel.type}})
{%- endfor %}    
''',
}


class EntityMeta(object):
    def __init__(self):
        pass

    def entity_abi(self, ent, format='json', target_dir=None):
        """
        $ python -m entity_meta entity_abi Person
        $ python -m entity_meta entity_abi Person simple
        $ python -m entity_meta entity_abi Person json ~/sagas/projs/hubs-common/asset/meta

        :param ent:
        :return:
        """
        from meta_generator import get_entity_abi
        from template_helper import render
        from os.path import expanduser

        meta = get_entity_abi(ent)

        if format == 'json':
            cnt = json.dumps(meta, indent=2, ensure_ascii=False)
            print(cnt)

            if target_dir is not None:
                target_dir = expanduser(target_dir)
                io_utils.write_to_file(join(target_dir, ent+".json"), cnt)
        else:
            # template = Template(simple_tpl)
            # cnt=template.render(ent=meta)
            cnt = render(templates[format], ent=meta)
            print(cnt)

    def write_all_abi(self, target_dir):
        """
        $ python -m entity_meta write_all_abi ~/sagas/projs/hubs-common/asset/meta
        :param target_dir:
        :return:
        """
        from sagas.ofbiz.entities import all_entities
        for ent in all_entities(False):
            self.entity_abi(ent, 'json', target_dir=target_dir)

    def all_entities(self):
        """
        $ python -m entity_meta all_entities
        :return:
        """
        from sagas.ofbiz.entities import all_entities
        all_ents=all_entities(False)
        print(f"total {len(all_ents)}")

if __name__ == '__main__':
    import fire

    fire.Fire(EntityMeta)
