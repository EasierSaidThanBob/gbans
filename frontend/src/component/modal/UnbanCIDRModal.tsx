import { useCallback } from 'react';
import NiceModal, { muiDialogV5, useModal } from '@ebay/nice-modal-react';
import {
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle
} from '@mui/material';
import Stack from '@mui/material/Stack';
import { Formik } from 'formik';
import { apiDeleteCIDRBan } from '../../api';
import { unbanValidationSchema } from '../../util/validators.ts';
import { UnbanReasonTextField } from '../formik/UnbanReasonTextField';
import { CancelButton, SubmitButton } from './Buttons';
import { UnbanFormValues, UnbanModalProps } from './UnbanSteamModal';

export const UnbanCIDRModal = NiceModal.create(({ banId }: UnbanModalProps) => {
    const modal = useModal();

    const onSubmit = useCallback(
        async (values: UnbanFormValues) => {
            if (values.unban_reason == '') {
                modal.reject({ error: 'Reason cannot be empty' });
                await modal.hide();
                return;
            }
            try {
                await apiDeleteCIDRBan(banId, values.unban_reason);
                modal.resolve();
            } catch (e) {
                modal.reject(e);
            } finally {
                await modal.hide();
            }
        },
        [banId, modal]
    );

    return (
        <Formik<UnbanFormValues>
            initialValues={{ unban_reason: '' }}
            onSubmit={onSubmit}
            validateOnChange={true}
            validateOnBlur={true}
            validationSchema={unbanValidationSchema}
        >
            <Dialog {...muiDialogV5(modal)}>
                <DialogTitle>Unban CIDR (#{banId})</DialogTitle>

                <DialogContent>
                    <Stack spacing={2}>
                        <UnbanReasonTextField />
                    </Stack>
                </DialogContent>

                <DialogActions>
                    <CancelButton />
                    <SubmitButton />
                </DialogActions>
            </Dialog>
        </Formik>
    );
});

export default UnbanCIDRModal;
