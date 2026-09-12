"""Verified Logic Reference for Future Engine/Server Porting"""
import random

def resolve_attack(base_accuracy, attacker_skill, target_defense, posture_bonus, weapon_min, weapon_max, skill_bonus, armor_mitigation_pct, resistance_mult, roll=None):
    hit_chance=max(0,min(100, base_accuracy+attacker_skill-target_defense-posture_bonus))
    if roll is None: roll=random.randint(1,100)
    hit=roll<=hit_chance
    dmg=None
    if hit:
        raw=random.randint(weapon_min, weapon_max)+skill_bonus
        dmg=raw*(1-armor_mitigation_pct/100.0)*resistance_mult
    return {"hit_chance":hit_chance,"roll":roll,"hit":hit,"damage":round(dmg,2) if dmg else 0}
