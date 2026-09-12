import random
from statistics import mean

def gen(resource_type, planet):
    stats=[max(0,min(1000, random.gauss(500,150))) for _ in range(9)]
    oq=mean(stats)
    return {"type":resource_type,"planet":planet,"oq":round(oq,2),"radius_km":random.randint(5,20)}
