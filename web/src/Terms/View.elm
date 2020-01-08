module Terms.View exposing (view)

import Element
    exposing
        ( Element
        , centerX
        , column
        , el
        , fill
        , fromRgb255
        , height
        , link
        , paddingEach
        , paragraph
        , px
        , text
        , width
        )
import Element.Font as Font
import Terms.Terms exposing (terms)


view : Element msg
view =
    column
        [ height fill
        , centerX
        , width <| px terms.container.width
        ]
        [ backLink, title, body ]


backLink : Element msg
backLink =
    el
        [ Font.size terms.backLink.size
        , paddingEach { padding | top = terms.backLink.paddingTop }
        , Font.color <| fromRgb255 terms.backLink.color
        ]
        (link []
            { url = "/"
            , label = text "← Home"
            }
        )


title : Element msg
title =
    el
        [ Font.size terms.title.size
        , Font.family [ Font.typeface terms.title.font ]
        , paddingEach { padding | top = terms.title.padding, bottom = terms.title.padding }
        ]
        (text "Terms of Use")


body : Element msg
body =
    paragraph
        [ Font.size terms.body.size
        , Font.family [ Font.typeface terms.body.font ]
        ]
        [ text termsTest ]


termsTest : String
termsTest =
    "This website is provided solely for your convenience and information. Pivotal Software, Inc. (Pivotal) does not endorse or make any representations about this website, and Pivotal is not responsible for the accuracy, reliability, and suitability of any information, data, opinions, advice or statements made on this website. Pivotal shall not be held liable for any losses or damages incurred by users of this website. Please note that your access to and usage of this website, including any materials, information, services and products described or provided therein, is made solely at your own risk and discretion."


padding : { top : Int, right : Int, bottom : Int, left : Int }
padding =
    { top = 0
    , right = 0
    , bottom = 0
    , left = 0
    }
