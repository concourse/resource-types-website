module Main exposing (Model, Msg, buildErrorMessage, layout, main, resourceTypeDecoder, update, view)

import Banner.View as Banner exposing (view)
import Browser
import Common.Common exposing (ResourceType)
import Element exposing (Element, column, fill, text, width)
import Html exposing (Html)
import Http
import Json.Decode exposing (Decoder, Error(..), field, list, map3, string)
import ResourceList.View as ResourceList exposing (view)


type Msg
    = DataReceived (Result Http.Error (List ResourceType))


type alias Model =
    { resourceTypes : List ResourceType
    , errorMessage : Maybe String
    }


apiUrl : String
apiUrl =
    "http://localhost:5019/resourceTypes"


main : Program () Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = \_ -> Sub.none
        }


init : () -> ( Model, Cmd Msg )
init _ =
    ( { resourceTypes = []
      , errorMessage = Nothing
      }
    , httpCommand
    )


view : Model -> Html Msg
view model =
    Element.layout [] (layout model)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        DataReceived (Ok resourceTypes) ->
            ( { model
                | resourceTypes = resourceTypes
                , errorMessage = Nothing
              }
            , Cmd.none
            )

        DataReceived (Err httpError) ->
            ( { model
                | errorMessage = Just (buildErrorMessage httpError)
              }
            , Cmd.none
            )


httpCommand : Cmd Msg
httpCommand =
    Http.get
        { url = apiUrl
        , expect = Http.expectJson DataReceived (list resourceTypeDecoder)
        }


buildErrorMessage : Http.Error -> String
buildErrorMessage httpError =
    case httpError of
        Http.BadUrl message ->
            message

        Http.Timeout ->
            "Server is taking too long to respond. Please try again later."

        Http.NetworkError ->
            "Unable to reach server."

        Http.BadStatus statusCode ->
            "Request failed with status code: " ++ String.fromInt statusCode

        Http.BadBody message ->
            message


layout : Model -> Element msg
layout model =
    column
        [ width fill ]
        [ Banner.view
        , case model.errorMessage of
            Just message ->
                text message

            Nothing ->
                ResourceList.view model.resourceTypes
        ]



-- order of fields have to match the order of ResourceType type


resourceTypeDecoder : Decoder ResourceType
resourceTypeDecoder =
    map3 ResourceType
        (field "name" string)
        (field "url" string)
        (field "description" string)
