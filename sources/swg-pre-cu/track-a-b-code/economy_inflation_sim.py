def simulate(days,credits,faucet,sink):
    hist=[]
    c=credits
    for d in range(days):
      c+=faucet-sink
      hist.append(c)
    return hist
