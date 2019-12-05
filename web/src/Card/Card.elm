module Card.Card exposing (Card, Container, card, container)

import Card.Styles as Styles exposing (borderRadius, height, hoverShadow, shadow, spacing, width)
import Common.Common exposing (Shadow)


type alias Card =
    { container : Container }


type alias Container =
    { height : Int
    , width : Int
    , borderRadius : Int
    , shadow : Shadow
    , hoverShadow : Shadow
    , spacing : Int
    }


card : Card
card =
    { container = container }


container : Container
container =
    { height = Styles.height
    , width = Styles.width
    , borderRadius = Styles.borderRadius
    , shadow = Styles.shadow
    , hoverShadow = Styles.hoverShadow
    , spacing = Styles.spacing
    }
