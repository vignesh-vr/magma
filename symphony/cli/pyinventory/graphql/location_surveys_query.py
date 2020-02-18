#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .datetime_utils import fromisoformat

from .survey_question_type_enum import SurveyQuestionType


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


def enum_field(enum_type):
    def encode_enum(value):
        return value.value

    def decode_enum(t, value):
        return t(value)

    return field(
        metadata={
            "dataclasses_json": {
                "encoder": encode_enum,
                "decoder": partial(decode_enum, enum_type),
            }
        }
    )



@dataclass_json
@dataclass
class LocationSurveysQuery:
    __QUERY__ = """
    query LocationSurveysQuery($id: ID!) {
  location: node(id: $id) {
    ... on Location {
      surveys {
        id
        name
        completionTimestamp
        sourceFile {
          id
          fileName
          storeKey
        }
        surveyResponses {
          formName
          questionFormat
          questionText
          boolData
          emailData
          latitude
          longitude
          phoneData
          textData
          floatData
          intData
          dateData
        }
      }
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class LocationSurveysQueryData:
        @dataclass_json
        @dataclass
        class Node:
            @dataclass_json
            @dataclass
            class Survey:
                @dataclass_json
                @dataclass
                class File:
                    id: str
                    fileName: str
                    storeKey: Optional[str] = None

                @dataclass_json
                @dataclass
                class SurveyQuestion:
                    questionText: str
                    formName: Optional[str] = None
                    questionFormat: Optional[SurveyQuestionType] = None
                    boolData: Optional[bool] = None
                    emailData: Optional[str] = None
                    latitude: Optional[Number] = None
                    longitude: Optional[Number] = None
                    phoneData: Optional[str] = None
                    textData: Optional[str] = None
                    floatData: Optional[Number] = None
                    intData: Optional[int] = None
                    dateData: Optional[int] = None

                id: str
                name: str
                completionTimestamp: int
                surveyResponses: List[SurveyQuestion]
                sourceFile: Optional[File] = None

            surveys: List[Survey]

        location: Optional[Node] = None

    data: Optional[LocationSurveysQueryData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client, id: str):
        # fmt: off
        variables = {"id": id}
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data