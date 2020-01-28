module Footer.View exposing (view)

import Common.Common as Common exposing (center)
import Element
    exposing
        ( Attribute
        , Element
        , centerX
        , fill
        , fromRgb255
        , height
        , htmlAttribute
        , link
        , newTabLink
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
import Url.Builder as Url


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
            , showLinkExternal contribute
            , text "|"
            , showLinkExternal feedback
            ]
        ]


showLink : Link -> Element msg
showLink linkValue =
    link
        linkAttributes
        (linkDetails linkValue)


showLinkExternal : Link -> Element msg
showLinkExternal linkValue =
    newTabLink
        linkAttributes
        (linkDetails linkValue)


linkAttributes : List (Attribute msg)
linkAttributes =
    [ htmlAttribute (Html.Attributes.class "footer-link") ]


linkDetails : Link -> { url : String, label : Element msg }
linkDetails linkValue =
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
    , url = "https://github.com/concourse/resource-types/blob/master/README.md"
    }


feedback : Link
feedback =
    { text = "Feedback"
    , url = "https://github.com/concourse/resource-types-website/issues/new/choose"
    }
