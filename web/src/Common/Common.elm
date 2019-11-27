module Common.Common exposing
    ( RGB
    , ResourceType
    , Shadow
    , center
    , darkGrey
    , greyishRed
    , white
    )

import Element exposing (Attribute, htmlAttribute)
import Html.Attributes exposing (style)



-- TYPES


type alias RGB =
    { red : Int
    , green : Int
    , blue : Int
    , alpha : Float
    }


type alias Shadow =
    { offsetX : Float
    , offsetY : Float
    , blur : Float
    , size : Float
    , color : RGB
    }


type alias ResourceType =
    { name : String
    , url : String
    , description : String
    }



-- STYLING


center : Attribute msg
center =
    htmlAttribute (Html.Attributes.style "margin" "0 auto")



-- COLORS


darkGrey : RGB
darkGrey =
    { red = 42, green = 50, blue = 57, alpha = 1 }


white : RGB
white =
    { red = 255, green = 255, blue = 255, alpha = 1 }


greyishRed : RGB
greyishRed =
    { red = 98, blue = 85, green = 85, alpha = 0.4 }
