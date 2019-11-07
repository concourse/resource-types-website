module MainTests exposing (..)

import Main exposing (view)
import Test exposing (Test, describe, test)
import Test.Html.Query as Query
import Test.Html.Selector exposing (class, style)


suite : Test
suite =
    describe "Main"
        [ test "is centered" <|
            \_ ->
                Main.view
                    |> Query.fromHtml
                    |> Query.has
                        [ style "margin" "0 auto" ]
        ]
