    <!-- Content Header (Page header) -->
    <section class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1>Run a workflow</h1>
          </div>
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item"><a href="#">Home</a></li>
              <li class="breadcrumb-item active">Run a Workflow</li>
            </ol>
          </div>
        </div>
      </div><!-- /.container-fluid -->
    </section>

    <!-- Main content -->
    <section class="content">
      <form method=post>
        <div class="row">
          <div class="col-12">
            <div class="card card-primary">
              <div class="card-header">
                <h3 class="card-title">{{ .WorkflowName }}</h3>

                <div class="card-tools">
                  <button type="button" class="btn btn-tool" data-card-widget="collapse" title="Collapse">
                    <i class="fas fa-minus"></i>
                  </button>
                </div>
              </div>
              <div class="card-body">
                <div class="form-group">
                  <label for="inputName">Target (could be a network, hostname or IP) </label>
                  <select class="target-select" name="targets[]" class="form-control" multiple="multiple"
                    style="width: 100%">
                  </select>
                </div>
                <div class="form-group">
                  <div class="custom-control custom-switch custom-switch-off-danger custom-switch-on-success">
                    <input type="checkbox" class="custom-control-input" id="customSwitch3" name="recurseOnSubdomain">
                    <label class="custom-control-label" for="customSwitch3">Recurse on all subdomain of the target in
                      the database. The target must be a domain</label>
                  </div>
                </div>
                <!-- /.card-body -->
              </div>
              <!-- /.card -->
            </div>
          </div>
        </div>
            <div class="card-footer">
              <a href="/recon/workflows/list" class="btn btn-secondary">Cancel</a>
              <input type="submit" value="Run" class="btn btn-success float-right">
            </div>
       
      </form>
    </section>
    <!-- /.content -->