export interface BaseForm {
  beforeSubmit?: () => void;
  afterSubmit?: () => void;
}
