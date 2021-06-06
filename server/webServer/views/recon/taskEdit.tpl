       <!-- Content Header (Page header) -->
       <section class="content-header">
         <div class="container-fluid">
           <div class="row mb-2">
             <div class="col-sm-6">
               <h1>Task Editor</h1>
             </div>
             <div class="col-sm-6">
               <ol class="breadcrumb float-sm-right">
                 <li class="breadcrumb-item"><a href="#">Home</a></li>
                 <li class="breadcrumb-item active">Text Editors</li>
               </ol>
             </div>
           </div>
         </div><!-- /.container-fluid -->
       </section>

       <!-- Main content -->
       <section class="content">
         <div class="row">
           <div class="col-md-12">
             <div class="card card-outline card-info">
               <div class="card-header">
                 <h3 class="card-title">
                   {{ .TaskName }}
                 </h3>
               </div>
               <!-- /.card-header -->
               <div class="card-body p-0">
                 <textarea id="codeMirrorDemo" class="p-3">
{{ .Yaml }}
              </textarea>
               </div>
             </div>
           </div>
           <!-- /.col-->
         </div>
         <div class="card-footer">
           <a href="{{urlfor "Recon.ListTasks"}}" class="btn btn-secondary">Cancel</a>
           <a href="{{urlfor "ReconController.RunTask" ":taskName"  .taskName }}" class="btn btn-success float-right">Run</a>
         </div>
         <!-- ./row -->
       </section>
       <!-- /.content -->