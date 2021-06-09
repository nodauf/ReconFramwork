 <!-- Main Sidebar Container -->
 <aside class="main-sidebar sidebar-dark-primary elevation-4">
   <!-- Brand Logo -->
   <a href="index3.html" class="brand-link">
     <img src="/static/dist/img/AdminLTELogo.png" alt="AdminLTE Logo" class="brand-image img-circle elevation-3"
       style="opacity: .8">
     <span class="brand-text font-weight-light">AdminLTE 3</span>
   </a>

   <!-- Sidebar -->
   <div class="sidebar">

     <!-- SidebarSearch Form -->
     <div class="form-inline">
       <div class="input-group" data-widget="sidebar-search">
         <input class="form-control form-control-sidebar" type="search" placeholder="Search" aria-label="Search">
         <div class="input-group-append">
           <button class="btn btn-sidebar">
             <i class="fas fa-search fa-fw"></i>
           </button>
         </div>
       </div>
     </div>

     <!-- Sidebar Menu -->
     <nav class="mt-2">
       <ul class="nav nav-pills nav-sidebar flex-column" data-widget="treeview" role="menu" data-accordion="false">
         <!-- Add icons to the links using the .nav-icon class
               with font-awesome or any other icon font library -->
         <li class="nav-item menu-open">
           <a href="#" class="nav-link active">
             <i class="nav-icon fas fa-tachometer-alt"></i>
             <p>
               Dashboard
             </p>
           </a>
         </li>
         <li class="nav-item">
           <a href="#" class="nav-link">
             <i class="nav-icon far fa-plus-square"></i>
             <p>
               Run tasks or workflows
             </p>
           </a>
         </li>
         <li class="nav-header">TASKS</li>

         <li class="nav-item">
           <a href="{{urlfor "ReconController.ListTasks"}}" class="nav-link">
             <i class="nav-icon fas fa-book"></i>
             <p>
               List the tasks
             </p>
           </a>

         </li>
         <li class="nav-header">WORKFLOWS</li>
         <li class="nav-item">
           <a href="{{urlfor "ReconController.ListWorkflows"}}" class="nav-link">
             <i class="nav-icon fas fa-book"></i>
             <p>
               List the workflows
             </p>
           </a>

         </li>
         <li class="nav-header">RESULTS</li>
         <li class="nav-item">
           <a href="{{urlfor "ReconController.OverviewResults"}}" class="nav-link">
             <i class="fas fa-circle nav-icon"></i>
             <p>Overview</p>
           </a>
         </li>
         <li class="nav-item">
           <a href="{{urlfor "ReconController.ListAllResults"}}" class="nav-link">
             <i class="fas fa-circle nav-icon"></i>
             <p>Details</p>
           </a>
         </li>
         <li class="nav-item">
           <a href="{{urlfor "ReconController.ListResultsWeb"}}" class="nav-link">
             <i class="nav-icon fas fa-circle"></i>
             <p>Web</p>
           </a>
         </li>
       </ul>
     </nav>
     <!-- /.sidebar-menu -->
   </div>
   <!-- /.sidebar -->
 </aside>