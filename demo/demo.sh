echo "KSL modules encapsulate the access management requirements for a subdomain (ex: a service)"
read -n 1 -s -r -p "Press any key to continue..."
echo
echo cat rbac.ksl
sleep 1
cat rbac.ksl
echo
read -n 1 -s -r -p "Press any key to continue..."
echo
echo "KSL files can be built directly into schemas for backend Zanzibar implementations."
read -n 1 -s -r -p "Press any key to continue..."
echo
echo ./ksl -o schema.zed rbac.ksl
sleep 1
./ksl -o schema.zed rbac.ksl
echo cat schema.zed
sleep 1
cat schema.zed
echo
read -n 1 -s -r -p "Press any key to continue..."
echo
echo "KSL files can also be compiled to an intermediate language"
read -n 1 -s -r -p "Press any key to continue..."
echo
echo ./ksl -c rbac.ksl -o rbac.json
sleep 1
./ksl -c rbac.ksl -o rbac.json
cat rbac.json
sleep 5
echo
echo "It's not as readable as KSL, but easy it's to serialize/deserialize and can be generated from other front-ends. Like TypeSpec!"
read -n 1 -s -r -p "Press any key to continue..."
echo
echo cat typespec/inventory.tsp
sleep 1
cat typespec/inventory.tsp
echo
sleep 5
echo "This API definition uses custom annotations to describe the required permissions, and we have custom emitters to output KSL IR like the following."
read -n 1 -s -r -p "Press any key to continue..."
echo
echo cat typespec/inventory.json
sleep 1
cat typespec/inventory.json
echo
read -n 1 -s -r -p "Press any key to continue..."
echo
echo "And we can generate a schema from any combination of ksl and compiled json files"
read -n 1 -s -r -p "Press any key to continue..."
echo
echo ./ksl -o combined.zed rbac.ksl typespec/inventory.json
sleep 1
./ksl -o combined.zed rbac.ksl typespec/inventory.json
cat combined.zed
echo
echo "Note the generated relations - the typespec code was able to call the extension from the ksl code!"