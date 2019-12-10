module ResourceList.Styles exposing
    ( maxWidth
    , outsideMargin
    , paddingVertical
    , spacing
    )

import Card.Styles as CardStyles
import Common.Common as Common


maxCards : Int
maxCards =
    4


maxWidth : Int
maxWidth =
    ((CardStyles.containerWidth + spacing) * maxCards) + (outsideMargin * 2)


paddingVertical : Int
paddingVertical =
    Common.gridSize * 10


spacing : Int
spacing =
    Common.gridSize * 2


outsideMargin : Int
outsideMargin =
    Common.gridSize * 10
