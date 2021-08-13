# ideas

## Workflow Engine

## Steps


- **WorkFlow** a workflow is just another step that can be embedded in another larger workflow
- **StartStep** always start a workflow with a start step
- **StopStep** you can optionally end a workflow with an end step
- **PipeStep** pipe just pipes a step to one or more steps
- **FactConditionalStep** checks a factsheet for a condition and sends to the next steps depending
- **FailStep** if a workflow gets to a FailStep it fails
- **WaitForItStep** it calls a func every step and will not proceed until fun returns either a true or an error or it crossed a timeout.
- **MoreInfoNeededStep**