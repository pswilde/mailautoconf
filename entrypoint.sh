#!/usr/bin/env bash
echo Removing old sample files…
rm /mailautoconf/config/*.sample.yaml

function write_file() {
  while read line;
  do
    first_char=${line:0:1}

    if [[ $first_char != "#" ]]; then
      line="#"$line
    fi
    echo $line >> $2
  done < $1
}

echo Setting up new sample config files…
def_conf="/mailautoconf/default-config/config.default.yaml"
new_conf="/mailautoconf/config/config.sample.yaml"
write_file $def_conf $new_conf

def_serv="/mailautoconf/default-config/services.default.yaml"
new_serv="/mailautoconf/config/services.sample.yaml"
write_file $def_serv $new_serv

echo New sample files copied

cd /mailautoconf
exec /mailautoconf/mailautoconf
