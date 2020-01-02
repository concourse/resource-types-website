module Footer.View exposing (view)

import Common.Common as Common exposing (center)
import Element
    exposing
        ( Element
        , centerX
        , fill
        , fromRgb255
        , height
        , htmlAttribute
        , link
        , px
        , row
        , spacing
        , text
        , width
        )
import Element.Background as Background exposing (color)
import Element.Font as Font exposing (center, color, family, size, typeface)
import Footer.Footer exposing (footer)
import Html.Attributes exposing (class)
import Url.Builder as Url exposing (relative)


type alias Link =
    { text : String
    , url : String
    }


view : Element msg
view =
    let
        container =
            footer.container

        linkStyles =
            footer.link
    in
    row
        [ height <| px container.height
        , width <| fill
        , Background.color <| fromRgb255 container.backgroundColor
        , Common.center
        ]
        [ row
            [ centerX
            , spacing container.spacing
            , Font.color <| fromRgb255 linkStyles.color
            , Font.size linkStyles.size
            , Font.family [ Font.typeface linkStyles.font ]
            ]
            [ showLink terms
            , text "|"
            , showLink contribute
            , text "|"
            , showLink feedback
            ]
        ]


showLink : Link -> Element msg
showLink linkValue =
    link
        [ htmlAttribute (Html.Attributes.class "footer-link") ]
        { url = linkValue.url
        , label = text linkValue.text
        }


terms : Link
terms =
    { text = "Terms of Use"
    , url = Url.absolute [ "terms" ] []
    }


contribute : Link
contribute =
    { text = "Contribute"
    , url = "http://example.com"
    }


feedback : Link
feedback =
    { text = "Feedback"
    , url = "http://example.com"
    }
