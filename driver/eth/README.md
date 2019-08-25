Eth binding
```bash
# Compile solidity code
solcjs --abi mesh.sol

# Generate GO binding
abigen --abi mesh_sol_Mesh.abi --pkg eth --out mesh.go
```