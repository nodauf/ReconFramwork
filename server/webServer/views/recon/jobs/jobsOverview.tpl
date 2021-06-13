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
                <th style="width: 15%">
                  Target
                </th>
                <th style="width: 10%">
                  Machinery task
                </th>
                <th style="width: 30%">
                  Machinery Arguments
                </th>
                <th style="width: 20%">
                  Execution date
                </th>
                <th style="width: 5%">
                Processed
                </th>
                <th style="width: 15%">
                Actions
                </th>
              </tr>
            </thead>
            <tbody>
              {{ range $result := .Results}}
              <tr>
                <td>
                  {{ $result.Target }}
                </td>
                <td>
                  {{ $result.MachineryTask }}
                </td>
                <td>
                  {{ $result.MachineryArgs }}
                </td>
                <td>
                  {{ $result.ExecutionTime }}
                </td>
                <td>
                  {{ if $result.Processed }}
                  Yes
                  {{ else }}
                  No
                  {{ end }}
                </td>
                <td class="text-right">
                  <a class="btn btn-info btn-sm" href="{{urlfor "ReconController.DetailsJob" ":id" $result.ID }}" data-toggle="modal" data-target="#exampleModal" data-remote="false">
                    <i class="fas fa-list">
                    </i>
                    Details
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