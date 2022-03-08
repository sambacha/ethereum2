#!/usr/bin/python3
import matplotlib.pyplot as plt
import matplotlib.ticker as mtick
import numpy as np
import re
import sys

attnets_size = 64

if len(sys.argv) < 2:
    print('the log file is missing in the argument')
    exit(1)

log_file = open(sys.argv[1], 'r')
lines = log_file.readlines()
counts = [0] * attnets_size

for line in lines:
    m = re.search(r"attnets=([0-9a-f]*)", line)
    if m == None:
        continue
    encoded_bytes = bytes.fromhex(m.group(1))

    for i in range(attnets_size):
        # follow the spec at https://github.com/ethereum/consensus-specs/blob/dev/ssz/simple-serialize.md#bitvectorn
        if (encoded_bytes[i // 8] & (1 << (i % 8))) != 0:
            counts[i] += 1

fig, ax = plt.subplots()
ax.bar(range(attnets_size), np.array(counts)/sum(counts) * 100)
ax.set_title('Distribution of numbers of nodes to attnets they are subscribed to')
ax.set_xlabel('subnet id from a total of 64 subnets')
ax.set_ylabel('percent of nodes subscribed to the subnet')
ax.yaxis.set_major_formatter(mtick.PercentFormatter())
plt.show()
