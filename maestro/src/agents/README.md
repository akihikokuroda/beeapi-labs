# Maestro agents

## Supported agents

The following agents are supported in Maestro

* BeeAI
* CrewAI
* OpenAI
* Remote
* Custom

### [BeeAI](https://docs.beeai.dev/introduction/welcome)

#### Configuration
### [CrewAI](https://www.crewai.com/)
#### Configuration
### [OpenAI](https://openai.com/)
#### Configuration
### Remote
#### Configuration
### Custom
#### Configuration

## Agent Definition example

* You can define your Agent in a declarative way using a YAML file, where you can use the current BeeAI Agent implementation. With that, you can configure your agent or agents. For example, create an `agents.yaml` file containing the following:

```yaml
apiVersion: maestro/v1alpha1
kind: Agent
metadata:
  name: current-affairs
  labels:
    app: mas-example
spec:
  model: meta-llama/llama-3-1-70b-instruct
  description: Get the current weather
  tools:
    - code_interpreter
    - weather
  instructions: Get the current temperature for the location provided by the user. Return results in Fahrenheit.

---
apiVersion: maestro/v1alpha1
kind: Agent
metadata:
  name: hot-or-not
  labels:
    app: mas-example
spec:
  model: meta-llama/llama-3-1-70b-instruct
  description: Is the current temperature hotter than usual?
  tools:
    - code_interpreter
    - weather
  instructions: The user will give you a temperature in Fahrenheit and a location. Use the OpenMateo weather tool to find the average monthly temperature for the location. Answer if the temperature provided by the user is hotter or colder than the average found by the tool.
```
