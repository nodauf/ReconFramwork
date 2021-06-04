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
               Run a task
               <i class="fas fa-angle-left right"></i>
             </p>
           </a>
           <ul class="nav nav-treeview">
             <li class="nav-item">
               <a href="#" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>
                   Login & Register v1
                   <i class="fas fa-angle-left right"></i>
                 </p>
               </a>
               <ul class="nav nav-treeview">
                 <li class="nav-item">
                   <a href="pages/examples/login.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Login v1</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="pages/examples/register.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Register v1</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="pages/examples/forgot-password.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Forgot Password v1</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="pages/examples/recover-password.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Recover Password v1</p>
                   </a>
                 </li>
               </ul>
             </li>
             <li class="nav-item">
               <a href="#" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>
                   Login & Register v2
                   <i class="fas fa-angle-left right"></i>
                 </p>
               </a>
               <ul class="nav nav-treeview">
                 <li class="nav-item">
                   <a href="pages/examples/login-v2.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Login v2</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="pages/examples/register-v2.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Register v2</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="pages/examples/forgot-password-v2.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Forgot Password v2</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="pages/examples/recover-password-v2.html" class="nav-link">
                     <i class="far fa-circle nav-icon"></i>
                     <p>Recover Password v2</p>
                   </a>
                 </li>
               </ul>
             </li>
             <li class="nav-item">
               <a href="pages/examples/lockscreen.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Lockscreen</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="pages/examples/legacy-user-menu.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Legacy User Menu</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="pages/examples/language-menu.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Language Menu</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="pages/examples/404.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Error 404</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="pages/examples/500.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Error 500</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="pages/examples/pace.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Pace</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="pages/examples/blank.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Blank Page</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="starter.html" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Starter Page</p>
               </a>
             </li>
           </ul>
         </li>
         <li class="nav-header">TASKS</li>

         <li class="nav-item">
           <a href="/recon/tasks/list" class="nav-link">
             <i class="nav-icon fas fa-book"></i>
             <p>
               List the tasks
             </p>
           </a>

         </li>
         <li class="nav-header">WORKFLOWS</li>
         <li class="nav-item">
           <a href="/recon/workflows/list" class="nav-link">
             <i class="nav-icon fas fa-book"></i>
             <p>
               List the workflows
             </p>
           </a>

         </li>
         <li class="nav-header">RESULTS</li>
         <li class="nav-item">
           <a href="#" class="nav-link">
             <i class="fas fa-circle nav-icon"></i>
             <p>Level 1</p>
           </a>
         </li>
         <li class="nav-item">
           <a href="#" class="nav-link">
             <i class="nav-icon fas fa-circle"></i>
             <p>
               Level 1
               <i class="right fas fa-angle-left"></i>
             </p>
           </a>
           <ul class="nav nav-treeview">
             <li class="nav-item">
               <a href="#" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Level 2</p>
               </a>
             </li>
             <li class="nav-item">
               <a href="#" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>
                   Level 2
                   <i class="right fas fa-angle-left"></i>
                 </p>
               </a>
               <ul class="nav nav-treeview">
                 <li class="nav-item">
                   <a href="#" class="nav-link">
                     <i class="far fa-dot-circle nav-icon"></i>
                     <p>Level 3</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="#" class="nav-link">
                     <i class="far fa-dot-circle nav-icon"></i>
                     <p>Level 3</p>
                   </a>
                 </li>
                 <li class="nav-item">
                   <a href="#" class="nav-link">
                     <i class="far fa-dot-circle nav-icon"></i>
                     <p>Level 3</p>
                   </a>
                 </li>
               </ul>
             </li>
             <li class="nav-item">
               <a href="#" class="nav-link">
                 <i class="far fa-circle nav-icon"></i>
                 <p>Level 2</p>
               </a>
             </li>
           </ul>
         </li>
         <li class="nav-item">
           <a href="#" class="nav-link">
             <i class="fas fa-circle nav-icon"></i>
             <p>Level 1</p>
           </a>
         </li>
       </ul>
     </nav>
     <!-- /.sidebar-menu -->
   </div>
   <!-- /.sidebar -->
 </aside>