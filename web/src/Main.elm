module Main exposing (Model, Msg(..), Page(..), buildErrorMessage, main, resourceTypeDecoder, update, view)

import Banner.View as Banner exposing (view)
import Browser exposing (UrlRequest(..))
import Browser.Navigation as Nav
import Common.Common exposing (ResourceType)
import Element exposing (Element, centerX, centerY, column, el, fill, height, html, text, width)
import Footer.View as Footer exposing (view)
import Html exposing (a, div, img, nav)
import Html.Attributes exposing (class, href, src)
import Http
import Json.Decode as Decode exposing (Decoder, list, string)
import Json.Decode.Pipeline exposing (optional, required)
import List exposing (repeat)
import RemoteData exposing (WebData)
import ResourceList.View as ResourceList exposing (view)
import Terms.View as Terms exposing (view)
import Url
import Url.Parser as Url exposing (Parser)


type Msg
    = ResourceTypesReceived (WebData (List ResourceType))
    | FetchResourceTypes
    | UrlChange Url.Url
    | LinkClicked UrlRequest


type Page
    = Index
    | Terms


type alias Model =
    { resourceTypes : WebData (List ResourceType)
    , flags : Flags
    , navKey : Nav.Key
    , page : Page
    }


type alias Flags =
    { githubIconImg : String
    , githubStarImg : String
    , bannerImg : String
    }


apiUrl : String
apiUrl =
    "/api/v1/resources"



--


main : Program Flags Model Msg
main =
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = \_ -> Sub.none
        , onUrlRequest = LinkClicked
        , onUrlChange = UrlChange
        }


init : Flags -> Url.Url -> Nav.Key -> ( Model, Cmd Msg )
init flags url key =
    ( { resourceTypes = RemoteData.Loading
      , flags = flags
      , navKey = key
      , page = urlToPage url
      }
    , fetchResourceTypes
    )


view : Model -> Browser.Document Msg
view model =
    { title = "Concourse Resource Types"
    , body =
        [ Element.layout []
            (column
                [ width fill
                , height fill
                ]
                (html navigation
                    :: (case model.page of
                            Index ->
                                viewResourceTypes model

                            Terms ->
                                viewTerms
                       )
                )
            )
        ]
    }


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        FetchResourceTypes ->
            ( { model | resourceTypes = RemoteData.Loading }, fetchResourceTypes )

        ResourceTypesReceived response ->
            ( { model | resourceTypes = response }
            , Cmd.none
            )

        LinkClicked urlRequest ->
            case urlRequest of
                Internal url ->
                    ( model, Nav.pushUrl model.navKey (Url.toString url) )

                External url ->
                    ( model, Nav.load url )

        UrlChange url ->
            ( { model | page = urlToPage url }
            , Cmd.none
            )


urlToPage : Url.Url -> Page
urlToPage url =
    url
        |> Url.parse urlParser
        |> Maybe.withDefault Index


urlParser : Parser (Page -> a) a
urlParser =
    Url.oneOf
        [ Url.map Index Url.top
        , Url.map Terms (Url.s "terms")
        ]



--


fetchResourceTypes : Cmd Msg
fetchResourceTypes =
    Http.get
        { url = apiUrl
        , expect =
            list resourceTypeDecoder
                |> Http.expectJson (RemoteData.fromResult >> ResourceTypesReceived)
        }


viewResourceTypes : Model -> List (Element msg)
viewResourceTypes model =
    [ Banner.view model.flags.bannerImg
    , case model.resourceTypes of
        RemoteData.NotAsked ->
            el textStyles (text "")

        RemoteData.Loading ->
            el textStyles spinner

        RemoteData.Success resourceTypes ->
            ResourceList.view resourceTypes model.flags.githubIconImg model.flags.githubStarImg

        RemoteData.Failure httpError ->
            el textStyles (text <| buildErrorMessage httpError)
    , Footer.view
    ]


viewTerms : List (Element msg)
viewTerms =
    [ Terms.view, Footer.view ]


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
    [ centerX, centerY ]


spinner : Element msg
spinner =
    html
        (Html.div [ Html.Attributes.class "la-line-spin-clockwise-fade-rotating la-dark la-2x" ]
            (repeat 8 (Html.div [] []))
        )



-- order of fields have to match the order of ResourceType type


resourceTypeDecoder : Decoder ResourceType
resourceTypeDecoder =
    Decode.succeed ResourceType
        |> required "name" string
        |> required "repo" string
        |> optional "description" string ""
        |> required "username" string
        |> required "stars" string



-- top navigation from the concourse docs site


navigation : Html.Html msg
navigation =
    let
        baseUrl =
            "https://concourse-ci.org"

        blogUrl =
            "https://blog.concourse-ci.org"

        discussUrl =
            "https://discuss.concourse-ci.org"
    in
    div [ class "page-top" ]
        [ nav
            [ class "top-nav" ]
            [ a [ href <| baseUrl, class "top-link logo-link" ]
                [ img
                    [ src <| baseUrl ++ "/images/logo-white.svg" ]
                    []
                , Html.text "Concourse"
                ]
            , a [ href <| baseUrl ++ "/docs.html", class "top-link" ] [ Html.text "docs" ]
            , a [ href <| baseUrl ++ "/examples.html", class "top-link" ] [ Html.text "examples" ]
            , a [ href <| baseUrl ++ "/project.html", class "top-link" ] [ Html.text "project" ]
            , a [ href <| blogUrl, class "top-link" ] [ Html.text "blog" ]
            , a [ href <| discussUrl, class "top-link" ] [ Html.text "discuss" ]
            , a [ href "/", class "top-link active" ] [ Html.text "resource types" ]
            ]
        ]
