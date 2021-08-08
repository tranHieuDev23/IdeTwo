import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: ':id',
    loadChildren: () =>
      import('./pages/code-page/code-page.module').then(
        (m) => m.CodePageModule
      ),
  },
  {
    path: '**',
    loadChildren: () =>
      import('./pages/code-page/code-page.module').then(
        (m) => m.CodePageModule
      ),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
