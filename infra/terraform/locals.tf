locals {
  function_name = "shipping-package-size-calculator"
  function_description = "Shipping Package Size Calculator Function"
  function_memory = 512
  src_path      = "../../shipping_package_size_calculator/cmd"

  function_env_package_sizes = "[250, 500, 1000, 2000, 5000]"

  binary_name  = local.function_name
  binary_path  = "${path.module}/tf_generated/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.function_name}.zip"
}

output "binary_path" {
  value = local.binary_path
}