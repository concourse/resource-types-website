module Main exposing (layout, main)

import Banner.View as Banner exposing (view)
import Browser
import Element exposing (Element, column, fill, width, text)
import Html exposing (Html)
import Http
import Json.Decode exposing (Decoder, Error(..), decodeString, field, list, map3, string)
import ResourceList.View as ResourceList exposing (view)
import Common.Common exposing (ResourceType)


main : Program () Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = \_ -> Sub.none
        }

--main : Html msg
--main =
--    Element.layout [] layout

type Msg
    = SendHttpRequest
    | DataReceived (Result Http.Error (List ResourceType))


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        SendHttpRequest ->
            ( model, httpCommand )

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


init : () -> ( Model, Cmd Msg )
init _ =
    ( { resourceTypes = []
      , errorMessage = Nothing
      }
    , httpCommand
    )


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

view : Model -> Html Msg
view model =
    Element.layout [] (layout model)




layout : Model ->Element msg
layout model  =
    column
        [ width fill ]
        [ Banner.view
        , ResourceList.view model.resourceTypes
        
        ]



-- good ol spikin


apiUrl : String
apiUrl =
    "http://localhost:5019/resourceTypes"



type alias Model =
    { resourceTypes : List ResourceType
    , errorMessage : Maybe String
    }

resourceTypeDecoder : Decoder ResourceType
resourceTypeDecoder = map3 ResourceType 
  (field "name" string)
  (field "description" string)
  (field "url" string)

httpCommand : Cmd Msg
httpCommand =
    Http.get
        { url = apiUrl
        , expect = Http.expectJson DataReceived (list resourceTypeDecoder)
        }
