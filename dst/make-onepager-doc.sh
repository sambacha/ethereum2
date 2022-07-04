#!/usr/bin/env bash

#for f in *.json; do echo -e "# ------ $f ------"; cat $f; echo -e "\n"; done  > TMP_JSON.txt

echo "Generating Onepager..."
for f in *.json; do echo -e "### $f" echo -e "\n\`\`\`json"; cat $f; echo -e "\n\`\`\`" ; done  > 1-jsonrpc.md
echo "Done.."
sleep 1

exit 0



