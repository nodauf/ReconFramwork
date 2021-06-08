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
                <th style="width: 35%">
                  IP
                </th>
                <th style="width: 40%">
                  Domain
                </th>
                <th style="width: 15%">
                  Number of port open
                </th>
                <th style="width: 10%">
                </th>
              </tr>
            </thead>
            <tbody>
              {{ range $result := .Results}}
              <tr>
                <td>
                  {{ $result.Address }}
                </td>
                <td>
                  {{ range $index, $domain := $result.Domain}}
                  {{if $index}},{{end}}
                  {{ $domain }}
                  {{ end }}
                </td>
                <td>
                  {{ $result.NbPorts }}
                </td>
                <td class="text-right">
                  <a class="btn btn-info btn-sm" href="{{urlfor "ReconController.TreeResults" ":ip" $result.Address}}">
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