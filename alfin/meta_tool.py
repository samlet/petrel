from sagas.modules.deles import *

class MetaTool(object):
    def search_ent(self, filter, show_desc=False, show_auto_fields=True, show_view=False):
        """
        $ python meta_tool.py search_ent shopping
        $ python meta_tool.py search_ent shopping True False

        :param filter:
        :return:
        """
        from sagas.ofbiz.entities import all_entities, search_entity, entity
        result=search_entity(filter)
        for r in result:
            entity_model=entity(r)
            if not show_view and entity_model.is_view():
                continue

            print("-", f"{r} ({entity_model.package_name})")
            desc=entity_model.description
            if desc is not None:
                print("\t", desc)
            if show_desc:
                entity_model.desc(show_auto_fields)

    def list_package_ents(self, pkg="com.bluecc.workload"):
        """
        $ python meta_tool.py list_package_ents org.apache.ofbiz.product
        :param pkg:
        :return:
        """
        reader=oc.delegator.getModelReader()
        filters=oc.jset(pkg)
        ent_map=reader.getEntitiesByPackage(filters, oc.jset())
        total=0
        for k,v in ent_map.items():
            print(k)
            total=total+len(v)
            for ent in v:
                print("\t-", ent)
        print(f"total {total} entities in {len(ent_map)} packages")

if __name__ == '__main__':
    import fire
    fire.Fire(MetaTool)

