module Main exposing (Msg(..), buildErrorMessage, layout, main, resourceTypeDecoder, update, view)

import Banner.View as Banner exposing (view)
import Browser
import Common.Common exposing (Flags, Model, ResourceType, gridSize)
import Element exposing (Element, centerX, column, el, fill, html, padding, text, width)
import Html exposing (Html)
import Html.Attributes exposing (class)
import Http
import Json.Decode as Decode exposing (Decoder, list, string)
import Json.Decode.Pipeline exposing (optional, required)
import RemoteData exposing (WebData)
import ResourceList.View as ResourceList exposing (view)


type Msg
    = ResourceTypesReceived (WebData (List ResourceType))
    | FetchResourceTypes


apiUrl : String
apiUrl =
    "http://localhost:5019/resourceTypes"



--


main : Program Flags Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = \_ -> Sub.none
        }


init : Flags -> ( Model, Cmd Msg )
init flags =
    ( { resourceTypes = RemoteData.Loading
      , flags = flags
      }
    , fetchResourceTypes
    )


view : Model -> Html Msg
view model =
    Element.layout [] (layout model)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        FetchResourceTypes ->
            ( { model | resourceTypes = RemoteData.Loading }, fetchResourceTypes )

        ResourceTypesReceived response ->
            ( { model | resourceTypes = response }
            , Cmd.none
            )



--


fetchResourceTypes : Cmd Msg
fetchResourceTypes =
    Http.get
        { url = apiUrl
        , expect =
            list resourceTypeDecoder
                |> Http.expectJson (RemoteData.fromResult >> ResourceTypesReceived)
        }


layout : Model -> Element msg
layout model =
    column
        [ width fill ]
        (viewResourceTypes model)


viewResourceTypes : Model -> List (Element msg)
viewResourceTypes model =
    [ Banner.view model.flags.bannerImg
    , case model.resourceTypes of
        RemoteData.NotAsked ->
            el textStyles (text "")

        RemoteData.Loading ->
            el textStyles spinner

        RemoteData.Success resourceTypes ->
            ResourceList.view resourceTypes model.flags.githubIconImg

        RemoteData.Failure httpError ->
            el textStyles (text <| buildErrorMessage httpError)
    ]


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


textStyles : List (Element.Attribute msg)
textStyles =
    [ centerX
    , padding (gridSize * 10)
    ]


spinner : Element msg
spinner =
    html
        (Html.div [ Html.Attributes.class "la-line-spin-clockwise-fade-rotating la-dark la-2x" ]
            [ Html.div [] []
            , Html.div [] []
            , Html.div [] []
            , Html.div [] []
            , Html.div [] []
            , Html.div [] []
            , Html.div [] []
            , Html.div [] []
            ]
        )



-- order of fields have to match the order of ResourceType type


resourceTypeDecoder : Decoder ResourceType
resourceTypeDecoder =
    Decode.succeed ResourceType
        |> required "name" string
        |> required "url" string
        |> optional "description" string ""
