module ResourceList.ResourceListTests exposing (..)

import Expect exposing (equal)
import ResourceList.ResourceList as ResourceList exposing (container)
import ResourceList.Styles as Styles exposing (..)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Resource List Component"
        [ describe "container"
            [ test "has a max width" <|
                \_ ->
                    resourceListContainer.maxWidth
                        |> Expect.equal Styles.maxWidth
            , test "has vertical padding" <|
                \_ ->
                    resourceListContainer.paddingVertical
                        |> Expect.equal Styles.paddingVertical
            , test "has spacing for children" <|
                \_ ->
                    resourceListContainer.spacing
                        |> Expect.equal Styles.spacing
            ]
        ]


resourceListContainer =
    ResourceList.container
