  <script src="/static/plugins/GoJS/js/go.js"></script>
  <script src="/static/plugins/codemirror/codemirror.js"></script>
  <script src="/static/plugins/codemirror/addon/display/autorefresh.js"></script>
  <script id="code">
  function openModal(url){
  $("#exampleModal").modal('show');
  $(".modal-body").load(url, function () {
        $(function () {
            cm = CodeMirror.fromTextArea(document.getElementById("codeMirrorDemo"), {
                theme: "monokai",
                readOnly: true,
                autorefresh: true,
            });

            setTimeout(function () {
                cm.refresh()
            }, 200);

        })

    });
  }
    function init() {
      var $ = go.GraphObject.make; // for conciseness in defining templates
      myDiagram =
        $(go.Diagram, "myDiagramDiv", {
          allowCopy: false,
          "draggingTool.dragsTree": true,
          "commandHandler.deletesTree": true,
          layout: $(go.TreeLayout, {
            angle: 0,
            arrangement: go.TreeLayout.ArrangementFixedRoots
          }),
          "undoManager.isEnabled": true
        });

      // when the document is modified, add a "*" to the title and enable the "Save" button
      myDiagram.addDiagramListener("Modified", function (e) {
        var button = document.getElementById("SaveButton");
        if (button) button.disabled = !myDiagram.isModified;
        var idx = document.title.indexOf("*");
        if (myDiagram.isModified) {
          if (idx < 0) document.title += "*";
        } else {
          if (idx >= 0) document.title = document.title.substr(0, idx);
        }
      });

      var bluegrad = $(go.Brush, "Linear", {
        0: "#C4ECFF",
        1: "#70D4FF"
      });
      var greengrad = $(go.Brush, "Linear", {
        0: "#B1E2A5",
        1: "#7AE060"
      });
      var yellowgrad = $(go.Brush, "Linear", {
        0: "#FEC901",
        1: "#FEA200"
      });
      var graygrad = $(go.Brush, "Linear", {
        0: "#F5F5F5",
        1: "#F1F1F1"
      });
      var lavgrad = $(go.Brush, "Linear", {
        0: "#EF9EFA",
        1: "#A570AD"
      });

      // each action is represented by a shape and some text
      var actionTemplate =
        $(go.Panel, "Horizontal",
          $(go.Shape, {
              width: 12,
              height: 12
            },
            new go.Binding("figure"),
            new go.Binding("fill")
          ),
          $(go.TextBlock, {
              font: "10pt Verdana, sans-serif"
            },
            new go.Binding("text")
          )
        );

      // each regular Node has body consisting of a title followed by a collapsible list of actions,
      // controlled by a PanelExpanderButton, with a TreeExpanderButton underneath the body
      myDiagram.nodeTemplate = // the default node template
        $(go.Node, "Vertical",
          new go.Binding("isTreeExpanded").makeTwoWay(), // remember the expansion state for
          new go.Binding("wasTreeExpanded").makeTwoWay(), //   when the model is re-loaded
          {
            selectionObjectName: "BODY"
          },
          // the main "BODY" consists of a RoundedRectangle surrounding nested Panels
          $(go.Panel, "Auto", {
              name: "BODY"
            },
            $(go.Shape, "RoundedRectangle", {
              fill: bluegrad,
              stroke: "white"
            }),
            $(go.Panel, "Vertical", {
                margin: 5
              },
              // the title
              $(go.TextBlock, {
                  stretch: go.GraphObject.Horizontal,
                  font: "bold 12pt Verdana, sans-serif",
                  margin: 5
                },
                new go.Binding("text", "value")
              ),
              // the optional list of actions
              $(go.Panel, "Vertical", {
                  stretch: go.GraphObject.Horizontal,
                  visible: false
                }, // not visible unless there is more than one action
                new go.Binding("visible", "actions", function (acts) {
                  return (Array.isArray(acts) && acts.length > 0);
                }),
                // headered by a label and a PanelExpanderButton inside a Table
                $(go.Panel, "Table", {
                    stretch: go.GraphObject.Horizontal
                  },
                  $(go.TextBlock, "Choices", {
                    alignment: go.Spot.Left,
                    font: "10pt Verdana, sans-serif"
                  }),
                  $("PanelExpanderButton", "COLLAPSIBLE", // name of the object to make visible or invisible
                    {
                      column: 1,
                      alignment: go.Spot.Right
                    }
                  )
                ), // end Table panel
                // with the list data bound in the Vertical Panel
                $(go.Panel, "Vertical", {
                    name: "COLLAPSIBLE", // identify to the PanelExpanderButton
                    padding: 2,
                    stretch: go.GraphObject.Horizontal, // take up whole available width
                    background: "white", // to distinguish from the node's body
                    defaultAlignment: go.Spot.Left, // thus no need to specify alignment on each element
                    itemTemplate: actionTemplate // the Panel created for each item in Panel.itemArray
                  },
                  new go.Binding("itemArray", "actions") // bind Panel.itemArray to nodedata.actions
                ) // end action list Vertical Panel
              ) // end optional Vertical Panel
            ) // end outer Vertical Panel
          ), // end "BODY"  Auto Panel
          $(go.Panel, // this is underneath the "BODY"
            {
              height: 17
            }, // always this height, even if the TreeExpanderButton is not visible
            $("TreeExpanderButton")
          )
        );

      // define a second kind of Node:
      myDiagram.nodeTemplateMap.add("Terminal",
        $(go.Node, "Spot",
          $(go.Shape, "Circle", {
            width: 55,
            height: 55,
            fill: greengrad,
            stroke: "white"
          }),
          $(go.TextBlock, {
              font: "10pt Verdana, sans-serif"
            },
            new go.Binding("text")
          )
        )
      );

      myDiagram.nodeTemplateMap.add("Port",
        $(go.Node, "Spot",
          $(go.Shape, "RoundedRectangle", {
            width: 55,
            height: 55,
            fill: yellowgrad,
            stroke: "white"
          }),
          $(go.TextBlock, {
              font: "10pt Verdana, sans-serif"
            },
            new go.Binding("text", "value")
          )
        )
      );

      myDiagram.nodeTemplateMap.add("PortComment",
        $(go.Node, "Spot",
          $(go.Shape, "RoundedRectangle", {
            height: 55,
            width: 120,
            fill: lavgrad,
            stroke: "black",
            click: function(e, obj) { openModal(obj.part.data.url) },
          }),
          $(go.TextBlock, {
              minSize: new go.Size(50, NaN),
              maxSize: new go.Size(80, NaN),
              margin: new go.Margin(0, 4, 0, 0),
              font: "10pt Verdana, sans-serif",
            click: function(e, obj) { openModal(obj.part.data.url) },
            },
            new go.Binding("text", "value")
          )
        )
      );

      myDiagram.nodeTemplateMap.add("Domain",
        $(go.Node, "Spot",
          $(go.Shape, "RoundedRectangle", {
            height: 55,
            width: 120,
            fill: graygrad,
            stroke: "black"
          }),
          $(go.TextBlock, {
              minSize: new go.Size(50, NaN),
              maxSize: new go.Size(80, NaN),
              margin: new go.Margin(0, 4, 0, 0),
              font: "10pt Verdana, sans-serif"
            },
            new go.Binding("text", "value")
          )
        )
      );

      myDiagram.linkTemplate =
        $(go.Link, go.Link.Orthogonal, {
            deletable: false,
            corner: 10
          },
          $(go.Shape, {
            strokeWidth: 2,
            stroke:"white"
          }),
          $(go.TextBlock, go.Link.OrientUpright, {
              background: "white",
              visible: false, // unless the binding sets it to true for a non-empty string
              segmentIndex: -2,
              segmentOrientation: go.Link.None
            },
            new go.Binding("text", "answer"),
            // hide empty string;
            // if the "answer" property is undefined, visible is false due to above default setting
            new go.Binding("visible", "answer", function (a) {
              return (a ? true : false);
            })
          )
        );
    
      {{ toJS . }}

      // create the Model with the above data, and assign to the Diagram
      myDiagram.model = $(go.GraphLinksModel, {
        copiesArrays: true,
        copiesArrayObjects: true,
        nodeDataArray: nodeDataArray,
        linkDataArray: linkDataArray
      });

    }
    window.addEventListener('DOMContentLoaded', init);
  </script>