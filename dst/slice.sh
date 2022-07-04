#!/bin/bash

SLICE_SIZE=2

JQ_SLICE_INPUTS='
   2376123525 as $EOF |            # random number that does not occur in the input stream to mark the end of the stream
   foreach (inputs, $EOF) as $input
   (
      # init state
      [[], []];                    # .[0]: array to collect inputs
                                   # .[1]: array that has collected $sliceSize inputs and is ready to be extracted
      # update state
      if .[0] | length == $sliceSize   # enough inputs collected
         or $input == $EOF             # or end of stream reached
      then [[$input], .[0]]        # create new array to collect next inputs. Save array .[0] with $sliceSize inputs for extraction
      else [.[0] + [$input], []]   # collect input, nothing to extract after this state update
      end;

      # extract from state
      if .[1] | length != 0
      then .[1]                    # extract array that has collected $sliceSize inputs
      else empty                   # nothing to extract right now (because still collecting inputs into .[0])
      end
   )
'

write_files() {
  local FILE_NAME_PREFIX=$1
  local FILE_COUNTER=0
  while read line; do
    FILE_COUNTER=$((FILE_COUNTER + 1))
    FILE_NAME="${FILE_NAME_PREFIX}_$FILE_COUNTER.json"
    echo "writing $FILE_NAME"
    jq '.'      > $FILE_NAME <<< "$line"   # array of formatted json inputs
#   jq -c '.'   > $FILE_NAME <<< "$line"   # compact array of json inputs
#   jq '.[]'    > $FILE_NAME <<< "$line"   # stream of formatted json inputs
#   jq -c '.[]' > $FILE_NAME <<< "$line"   # stream of compact json inputs
  done
}


echo "how to slice a stream of json inputs"
jq -n '{id: (range(5) + 1), a:[1,2]}' |   # create a stream of json inputs
jq -n -c --argjson sliceSize $SLICE_SIZE "$JQ_SLICE_INPUTS" |
write_files "stream_of_json_inputs_sliced"

echo -e "\nhow to slice an array of json inputs"
jq -n '[{id: (range(5) + 1), a:[1,2]}]' |                  # create an array of json inputs
jq -n --stream 'fromstream(1|truncate_stream(inputs))' |   # remove outer array to create stream of json inputs
jq -n -c --argjson sliceSize $SLICE_SIZE "$JQ_SLICE_INPUTS" |
write_files "array_of_json_inputs_sliced"