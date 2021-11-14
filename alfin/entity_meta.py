import json

simple_tpl='''
{{ent.name}}
fields:
{%- for fld in ent.fields %}
    - {{fld.name}} ({{fld.type}})  
{%- endfor %}

relations:
{%- for rel in ent.relations %}
    - {{rel.name}} ({{rel.type}})
{%- endfor %}    
'''

class EntityMeta(object):
    def __init__(self):
        pass

    def entity_abi(self, ent, format='json'):
        """
        $ python -m entity_meta entity_abi Person
        $ python -m entity_meta entity_abi Person tpl

        :param ent:
        :return:
        """
        from meta_generator import get_entity_abi
        from template_helper import render
        meta=get_entity_abi(ent)

        if format=='json':
            cnt=json.dumps(meta, indent=2, ensure_ascii=False)
            print(cnt)
        else:
            # template = Template(simple_tpl)
            # cnt=template.render(ent=meta)
            cnt=render(simple_tpl, ent=meta)
            print(cnt)

if __name__ == '__main__':
    import fire
    fire.Fire(EntityMeta)


