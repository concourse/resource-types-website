module ResourceList.View exposing (view)

import Card.View as CardView exposing (view)
import Common.Overrides as Overrides exposing (grid)
import Element exposing (Element, fill, maximum, paddingXY, width, wrappedRow)
import ResourceList.ResourceList exposing (container)


view : Element msg
view =
    wrappedRow
        ([ width (fill |> maximum container.maxWidth)
         , paddingXY container.outsideMargin 0
         ]
            ++ Overrides.grid
        )
        [ CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        , CardView.view
        ]
