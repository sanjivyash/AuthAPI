#!bin/bash
cd tests
bash script.sh > out 
cmp --silent out check && echo "All Tests Passed!!" || echo "Some Tests Failed"
cd ..