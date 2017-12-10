#!/bin/bash
errors=()
temp=$(date '+%Y-%m-%d_%T')
mkdir results/$temp
set -e

pattern="${TEST:-*}.scala"

for f in $(find user-files/simulations -name ${pattern}); do
  package=$(awk '/package/{print $NF}' $f)
  class=$(awk '/class .* extends Simulation/{print $2}' $f)
  if [[ -z "$class" ]]; then
    continue
  fi
  result=$(gatling.sh --mute --simulation ${package}.${class} --results-folder results/$temp 2> results/$temp/errors.log)
  echo "$result" >> results/$temp/$class.log
  if [[ -s "results/$temp/errors.log" ]]; then
    cat results/$temp/errors.log
    exit 1
  fi
  error=$(echo "$result" | awk '/---- Global Information -+/{y=1}y;/==+/{y=0}' | awk '/---- Errors -+/{y=1}y;/==+/{y=0}')
  if [[ -n "$error" ]] ; then
    errors+=("${package}.${class}")
    errors+=("${error}")
  fi
done

printf "%s\n" "${errors[@]}"

if [[ -n "$errors" ]] ; then
  exit 1
fi
