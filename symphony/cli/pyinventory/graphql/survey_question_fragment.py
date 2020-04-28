#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from gql.gql.enum_utils import enum_field
from .survey_question_type_enum import SurveyQuestionType

QUERY: List[str] = ["""
fragment SurveyQuestionFragment on SurveyQuestion {
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

"""]

@dataclass
class SurveyQuestionFragment(DataClassJsonMixin):
    questionText: str
    formName: Optional[str]
    questionFormat: Optional[SurveyQuestionType] = enum_field(SurveyQuestionType)
    boolData: Optional[bool]
    emailData: Optional[str]
    latitude: Optional[Number]
    longitude: Optional[Number]
    phoneData: Optional[str]
    textData: Optional[str]
    floatData: Optional[Number]
    intData: Optional[int]
    dateData: Optional[int]