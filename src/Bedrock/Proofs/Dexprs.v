Require Import Coq.Lists.List.
Require Import bedrock2.Syntax.
Require Import bedrock2.WeakestPreconditionProperties.
Require Import coqutil.Map.Interface.
Require Import coqutil.Word.Interface.
Require Import Crypto.Bedrock.Tactics.
Require Import Crypto.Util.ListUtil.
Import ListNotations.

(* Proofs about [WeakestPrecondition.dexprs] *)
(* TODO: add to bedrock2? *)
(* TODO: some of these may be unused *)
Section Dexprs.
  Context {p : Semantics.parameters} {ok : @Semantics.parameters_ok p}.

  Lemma dexprs_cons_iff m l :
    forall e es v vs,
      WeakestPrecondition.dexprs m l (e :: es) (v :: vs) <->
      (WeakestPrecondition.expr m l e (eq v)
       /\ WeakestPrecondition.dexprs m l es vs).
  Proof.
    cbv [WeakestPrecondition.dexprs].
    induction es; split; intros;
      repeat match goal with
             | _ => progress cbn [WeakestPrecondition.list_map
                                    WeakestPrecondition.list_map_body] in *
             | _ => progress cbv beta in *
             | H : _ :: _ = _ :: _ |- _ => inversion H; clear H; subst
             | |- _ /\ _ => split
             | _ => progress destruct_head'_and
             | _ => reflexivity
             | _ => solve [propers]
             | _ => peel_expr
             end.
    eapply IHes with (vs := tl vs).
    propers_step. peel_expr.
    destruct vs; cbn [tl]; propers.
  Qed.

  Lemma dexprs_cons_nil m l e es :
    WeakestPrecondition.dexprs m l (e :: es) [] -> False.
  Proof.
    cbv [WeakestPrecondition.dexprs].
    induction es; intros;
      repeat match goal with
             | _ => progress cbn [WeakestPrecondition.list_map
                                    WeakestPrecondition.list_map_body] in *
             | _ => congruence
             | _ => solve [propers]
             | _ => apply IHes
             | _ => peel_expr
             end.
    propers_step. peel_expr. propers.
  Qed.

  Lemma dexprs_app_l m l es1 :
    forall es2 vs,
      WeakestPrecondition.dexprs m l (es1 ++ es2) vs ->
      (WeakestPrecondition.dexprs m l es1 (firstn (length es1) vs)) /\
      (WeakestPrecondition.dexprs m l es2 (skipn (length es1) vs)).
  Proof.
    induction es1; intros.
    { cbn [Datatypes.length skipn firstn]; rewrite app_nil_l in *.
      split; eauto; reflexivity. }
    { destruct vs; rewrite <-app_comm_cons in *;
        [ match goal with H : _ |- _ => apply dexprs_cons_nil in H; tauto end | ].
      cbn [Datatypes.length skipn firstn].
      rewrite !dexprs_cons_iff in *.
      destruct_head'_and.
      repeat split; try eapply IHes1; eauto. }
  Qed.

  Lemma dexprs_length m l :
    forall vs es,
      WeakestPrecondition.dexprs m l es vs ->
      length es = length vs.
  Proof.
    induction vs; destruct es; intros;
      repeat match goal with
             | _ => progress cbn [Datatypes.length]
             | _ => progress destruct_head'_and
             | H : _ |- _ => apply dexprs_cons_nil in H; tauto
             | H : _ |- _ => apply dexprs_cons_iff in H
             | _ => reflexivity
             | _ => congruence
             end.
    rewrite IHvs; auto.
  Qed.

  Lemma list_map_app_iff {A B}
        (f : A -> (B -> Prop) -> Prop)
        (f_ext :
           forall a H1 H2,
             (forall b, H1 b <-> H2 b) ->
             f a H1 <-> f a H2)
        xs ys post :
    WeakestPrecondition.list_map f (xs ++ ys) post <->
    WeakestPrecondition.list_map
      f xs (fun xx =>
              WeakestPrecondition.list_map
                f ys (fun yy => post (xx ++ yy))).
  Proof.
    revert ys post; induction xs;
      repeat match goal with
             | _ => progress intros
             | _ => progress cbn [WeakestPrecondition.list_map
                                    WeakestPrecondition.list_map_body] in *
             | _ => rewrite app_nil_l
             | _ => rewrite <-app_comm_cons
             | |- f _ _ <-> f _ _ => apply f_ext
             | _ => reflexivity
             end.
    apply IHxs.
  Qed.

  Lemma dexpr_equiv m l n x1 x2 :
    WeakestPrecondition.dexpr m l (expr.var n) x1 ->
    WeakestPrecondition.dexpr m l (expr.var n) x2 ->
    x1 = x2.
  Proof.
    destruct 1; destruct 1; cleanup; congruence.
  Qed.

  Lemma dexpr_put_same m l n x :
    WeakestPrecondition.dexpr m (map.put l n x) (expr.var n) x.
  Proof. eexists; rewrite map.get_put_same; tauto. Qed.

  Lemma dexpr_put_diff m l n1 n2 x y :
    n1 <> n2 ->
    WeakestPrecondition.dexpr m l (expr.var n1) x ->
    WeakestPrecondition.dexpr m (map.put l n2 y) (expr.var n1) x.
  Proof.
    destruct 2; intros; eexists; rewrite map.get_put_diff; eauto.
  Qed.
End Dexprs.
