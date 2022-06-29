terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
}

data "coxedge_environments" "test" {

}

output "envs" {
  value = data.coxedge_environments.test
}

# VM Workloads
resource "coxedge_workload" "test" {
  name = "test2"
  environment_name = data.coxedge_environments.test.environments[0].name
  type = "VM"
  image = "stackpath-edge/ubuntu-1804-bionic:v202104291427"
  add_anycast_ip_address = true
  persistent_storage {
    path = "/lib/data/newFolder"
    size = 1
  }
  first_boot_ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDcYr9OnzsDfYVW2I1kX/iYJ0mPG490bI5mbxbOAKPLMuWLguxRohX804j1XbwZJ+Sna+9rSfxaYA8vgd1MoYX10l9cnMLx/MMbYp4ZquauN4pGY3WoDeCqsTss3VUMW+7RFBILpU3SJTlDV02FI36D3IXb4A8XymCyU3KC99XXTfTQsuKC+WFRMsTWtklrasqCVd5yEG90i/aJc6A3TZGOYgPFNEeVYvNDaJmIkb3y4FfShoBIMgZRt0ay7SvWZUvyfvyNmK5W9ePdhZZ58R+7tQNmCzjQ4v0suWRuGJ/XL3+03w3HEsDdQx+noL+R+qAjoNFwc0spBBhJK+Q4ADqr nothing@gmail.com"
  specs = "SP-1"
  deployment {
    name = "testvm"
    enable_autoscaling = false
    pops = ["PVD"]
    instances_per_pop = 1
  }
}
