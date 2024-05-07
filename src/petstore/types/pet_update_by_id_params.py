# File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

from __future__ import annotations

from typing_extensions import TypedDict

from typing import List, Union, Dict, Optional
from typing_extensions import Literal, TypedDict, Required, Annotated
from .._types import FileTypes
from .._utils import PropertyInfo
from ..types import shared_params

__all__ = ["PetUpdateByIDParams"]

class PetUpdateByIDParams(TypedDict, total=False):
    name: str
    """Name of pet that needs to be updated"""

    status: str
    """Status of pet that needs to be updated"""