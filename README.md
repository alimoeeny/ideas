# ideas
- **Concept** Concept represents a bit of knowledge that can be queried form the user or from the machines, an example concept is RBC count, a workflow can ask for RBC count from an algorithm or from the user
- **Idea** Idea represents an instant of one or more Concepts being realized, it's a map of concepts to measurments. For example when we are talking about someones RBC count at a specific time, basically we are associating the concept of RBC count with the measurement of their rbcs at a particular time
- **Measurement**
- **ConceptSet**
- **MutuallyExclusiveConceptSet**
- **IdeaSet**
-
#### Example




## Workflow Engine

## Steps


- **Workflow** a workflow is just another step that can be embedded in another larger workflow
- **StartStep** always start a workflow with a start step
- **StopStep** you can optionally end a workflow with an end step
- **PipeStep** pipe just pipes a step to one or more steps
- **FactConditionalStep** checks a factsheet for a condition and sends to the next steps depending
- **FailStep** if a workflow gets to a FailStep it fails
- **WaitForItStep** it calls a func every step and will not proceed until fun returns either a true or an error or it crossed a timeout.
- **MoreInfoNeededStep**