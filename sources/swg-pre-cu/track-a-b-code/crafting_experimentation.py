import random
OUTCOMES=[('critical_success',0.10,(0.05,0.10)),('success',0.50,(0.02,0.05)),('normal',0.30,(0.0,0.02)),('failure',0.08,(0,0)),('critical_failure',0.02,(-0.05,-0.02))]

def resource_stat(base_stat, oq, mult, exp_bonus):
    improved=base_stat+(oq*mult)
    return improved*(1+exp_bonus)
