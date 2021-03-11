module ResourceList.ResourceList exposing (Container, container)

import ResourceList.Styles as Styles exposing (maxWidth, outsideMargin, paddingVertical)


type alias Container =
    { maxWidth : Int
    , paddingVertical : Int
    , spacing : Int
    , outsideMargin : Int
    }


container : Container
container =
    { maxWidth = Styles.maxWidth
    , paddingVertical = Styles.paddingVertical
    , spacing = Styles.spacing
    , outsideMargin = Styles.outsideMargin
    }
