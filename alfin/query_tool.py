import records

class QueryTool(object):
    def __init__(self):
        self.db = records.Database('postgresql://xiaofeiwu@localhost/ofbiz')

    def status_items(self, status_type_id, format='json'):
        """
        # ORDER_RETURN_STTS, ORDER_STATUS, ORDER_ITEM_STATUS
        $ python -m query_tool status_items ORDER_STATUS | json
        $ python -m query_tool status_items ORDER_ITEM_STATUS | json

        :param fn:
        :param kwargs:
        :return:
        """
        rows = self.db.query('select * from status_item where status_type_id=:type_id',
                             type_id=status_type_id)
        print(rows.export(format))

    def status_types(self, format='json'):
        """
        $ python -m query_tool status_types
        $ python -m query_tool status_types csv

        :param fn:
        :param kwargs:
        :return:
        """
        rows = self.db.query('select DISTINCT status_type_id from status_item order by status_type_id')
        print(rows.export(format))

if __name__ == '__main__':
    import fire
    fire.Fire(QueryTool)
