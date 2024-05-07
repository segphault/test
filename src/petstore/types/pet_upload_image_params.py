# File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

from __future__ import annotations

from typing_extensions import TypedDict, Annotated

from .._utils import PropertyInfo

from typing import List, Union, Dict, Optional
from typing_extensions import Literal, TypedDict, Required, Annotated
from .._types import FileTypes
from .._utils import PropertyInfo
from ..types import shared_params

__all__ = ["PetUploadImageParams"]

class PetUploadImageParams(TypedDict, total=False):
    additional_metadata: Annotated[str, PropertyInfo(alias="additionalMetadata")]
    """Additional Metadata"""