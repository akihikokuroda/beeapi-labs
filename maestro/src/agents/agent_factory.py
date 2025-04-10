# SPDX-License-Identifier: Apache-2.0
from enum import StrEnum
from typing import Callable, Type, Union
import os, dotenv


from .beeai_agent import BeeAIAgent
from .beeai_local_agent import BeeAILocalAgent
from .crewai_agent import CrewAIAgent
from .remote_agent import RemoteAgent
from .mock_agent import MockAgent

EMOJIS = {
    'beeai': '🐝',
    'crewai': '👥',
    'mock': '🤖',
    'remote': '💸',

    # # Not yet supported
    # 'langflow': '⛓',
    # 'openai': '🔓',
}

dotenv.load_dotenv()

class AgentFramework(StrEnum):
    """Enumeration of supported frameworks"""
    BEEAI = "beeai"
    CREWAI = "crewai"
    MOCK = 'mock'
    REMOTE = 'remote'

    # Not yet supported
    # LANGFLOW = 'langflow'
    # OPENAI = 'openai'

class AgentFactory:
    """Factory class for handling agent frameworks"""
    @staticmethod
    def create_agent(framework: AgentFramework) -> Callable[..., Union[BeeAIAgent, CrewAIAgent]]:
        """Create an instance of the specified agent framework.

        Args:
            framework (AgentFramework): The framework to create. Must be a valid enum value.

        Returns:
            A new instance of the corresponding agent class.
        """
        local = os.getenv("LOCAL_AGENT", False)
        print(local)
        factories = {
            AgentFramework.BEEAI: BeeAILocalAgent,
            AgentFramework.CREWAI: CrewAIAgent,
            AgentFramework.MOCK: MockAgent
        }

        remote_factories = {
            AgentFramework.BEEAI: BeeAIAgent,
            AgentFramework.REMOTE: RemoteAgent,
            AgentFramework.MOCK: MockAgent
        }

        if framework not in factories:
            raise ValueError(f"Unknown framework: {framework}")

        if local:           
            return factories[framework]
        else:
            return remote_factories[framework]

    @classmethod
    def get_factory(cls, framework: str) -> Callable[..., Union[BeeAIAgent, BeeAILocalAgent, CrewAIAgent, RemoteAgent, MockAgent]]:
        """Get a factory function for the specified agent type."""
        print("!!!!!!")
        return cls.create_agent(framework)
