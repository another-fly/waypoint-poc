# Waypoint poc

## Description

From its [webpage](https://www.waypointproject.io):

> Waypoint provides a modern workflow to build, deploy, and release across platforms.
 
Waypoint uses a single configuration file and workflow to manage deployment across multiple platforms.

It has an [API SDK](https://github.com/hashicorp/waypoint-plugin-sdk) for developing. It has a MPL license.

## What we want to achieve:

- Abstract the deployment in multiple environments
- MVP with docker-compose
- CLI features
- SDK features

## Notes

- Waypoint cli uses HCL files to manage deployment
- Waypoint project should be created inside the project where the project/dockerfile is
  - The path can be specified, but not the port. It's always the 3000 tcp port
- Executing `waypoint init` will create an example `hcl` file. Set up the file and execute again `waypoint init` to create the project folders
- Waypoint with Hashicorp high coupling
