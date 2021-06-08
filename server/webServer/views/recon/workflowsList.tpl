    <!-- Content Header (Page header) -->
    <section class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1>Projects</h1>
          </div>
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item"><a href="#">Home</a></li>
              <li class="breadcrumb-item active">Projects</li>
            </ol>
          </div>
        </div>
      </div><!-- /.container-fluid -->
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">DataTable with minimal features & hover style</h3>
        </div>
        <!-- /.card-header -->
        <div class="card-body">
          <table id="tableHosts" class="table table-hover table-striped  ">
            <thead>
              <tr>
                <th style="width: 20%">
                  Task name
                </th>
                <th style="width: 30%">
                  Description
                </th>
                <th>
                  Tags
                </th>
                <th style="width: 20%">
                </th>
              </tr>
            </thead>
            <tbody>
              {{ range $workflowName, $workflow := .Workflows}}
              <tr>
                <td>
                  {{ $workflowName }}
                </td>
                <td>
                  {{ $workflow.Description }}
                </td>
                <td>
                  {{ range $index, $tag := $workflow.Tags}}
                  {{if $index}},{{end}}
                  {{ $tag }}
                  {{ end }}
                </td>
                <td class="project-actions text-right">
                  <a class="btn btn-success btn-sm" href="{{urlfor "ReconController.RunWorkflow" ":workflowName"  $workflowName }}">
                    <i class="fas fa-play">
                    </i>
                    Run
                  </a>
                  <a class="btn btn-info btn-sm" href="{{urlfor "ReconController.EditWorkflow" ":workflowName"  $workflowName }}">
                    <i class="fas fa-pencil-alt">
                    </i>
                    Edit
                  </a>
                  <a class="btn btn-danger btn-sm" href="/recon/workflows/delete/{{ $workflowName }}">
                    <i class="fas fa-trash">
                    </i>
                    Delete
                  </a>
                </td>
              </tr>
              {{ end }}
            </tbody>
          </table>
        </div>
        <!-- /.card-body -->
      </div>

    </section>
    <!-- /.content -->