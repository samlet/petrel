from jinja2 import environment
import stringcase
from jinja2 import Template, Environment

environment.DEFAULT_FILTERS['upper'] = lambda s: stringcase.uppercase(s)
environment.DEFAULT_FILTERS['lower'] = lambda s: stringcase.lowercase(s)

def render(tpl, **kwargs) -> str:
    template = Template(tpl)
    cnt = template.render(kwargs)
    return cnt

