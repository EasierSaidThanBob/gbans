import { useCallback, useMemo, useState } from 'react';
import NiceModal, { useModal } from '@ebay/nice-modal-react';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import ConstructionIcon from '@mui/icons-material/Construction';
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import EditIcon from '@mui/icons-material/Edit';
import LockIcon from '@mui/icons-material/Lock';
import Person2Icon from '@mui/icons-material/Person2';
import { Divider, IconButton, Theme } from '@mui/material';
import Avatar from '@mui/material/Avatar';
import Badge from '@mui/material/Badge';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import ButtonGroup from '@mui/material/ButtonGroup';
import Pagination from '@mui/material/Pagination';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';
import { useTheme } from '@mui/material/styles';
import { createFileRoute, useNavigate, useRouteContext } from '@tanstack/react-router';
import { isAfter } from 'date-fns/fp';
import { z } from 'zod';
import { PermissionLevel, permissionLevelString } from '../api';
import { apiDeleteMessage, ForumMessage, ForumThread } from '../api/forum.ts';
import { ForumRowLink } from '../component/ForumRowLink.tsx';
import { ForumThreadReplyBox } from '../component/ForumThreadReplyBox.tsx';
import { MarkDownRenderer } from '../component/MarkdownRenderer.tsx';
import RouterLink from '../component/RouterLink.tsx';
import { VCenterBox } from '../component/VCenterBox.tsx';
import { ModalConfirm, ModalForumThreadEditor } from '../component/modal';
import { useThread } from '../hooks/useThread.ts';
import { useThreadMessages } from '../hooks/useThreadMessages.ts';
import { logErr } from '../util/errors.ts';
import { useScrollToLocation } from '../util/history.ts';
import { commonTableSearchSchema, RowsPerPage } from '../util/table.ts';
import { avatarHashToURL, renderDateTime } from '../util/text.tsx';
import { LoginPage } from './_guest.login.index.tsx';

const forumThreadSearchSchema = z.object({
    ...commonTableSearchSchema
});

export const Route = createFileRoute('/_auth/forums/thread/$forum_thread_id')({
    component: ForumThreadPage,
    validateSearch: (search) => forumThreadSearchSchema.parse(search)
});

const ForumAvatar = ({ src, alt, online }: { src: string; alt: string; online: boolean }) => {
    const theme = useTheme();

    return (
        <Badge
            overlap={'circular'}
            anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
            variant="dot"
            sx={{
                '& .MuiBadge-badge': {
                    backgroundColor: online ? theme.palette.success.light : theme.palette.error.dark,
                    color: online ? theme.palette.success.light : theme.palette.error.dark,
                    boxShadow: `0 0 0 2px ${theme.palette.background.paper}`,
                    '&::after': {
                        position: 'absolute',
                        top: 0,
                        left: 0,
                        width: '100%',
                        height: '100%',
                        borderRadius: '50%',
                        animation: online ? 'ripple 1.2s infinite ease-in-out' : undefined,
                        border: '1px solid currentColor',
                        content: '""'
                    }
                },
                '@keyframes ripple': {
                    '0%': {
                        transform: 'scale(.8)',
                        opacity: 1
                    },
                    '100%': {
                        transform: 'scale(2.4)',
                        opacity: 0
                    }
                }
            }}
        >
            <Avatar variant={'circular'} sx={{ height: '120px', width: '120px' }} src={src} alt={alt} />
        </Badge>
    );
};

const ThreadMessageContainer = ({
    message,
    isFirstMessage,
    onDeleteSuccess
}: {
    message: ForumMessage;
    onDeleteSuccess: (forum_message_id: number) => void;
    isFirstMessage: boolean;
}) => {
    const [edit, setEdit] = useState(false);
    const [updatedMessage, setUpdatedMessage] = useState<ForumMessage>();
    const confirmModal = useModal(ModalConfirm);
    const { hasPermission, userSteamID } = useRouteContext({ from: '/_auth/forums/thread/$forum_thread_id' });
    const theme = useTheme();
    const navigate = useNavigate();

    const activeMessage = useMemo(() => {
        if (updatedMessage != undefined) {
            return updatedMessage;
        }
        return message;
    }, [message, updatedMessage]);

    const onUpdate = useCallback((updated: ForumMessage) => {
        setUpdatedMessage(updated);
        setEdit(false);
    }, []);

    const editable = useMemo(() => {
        return userSteamID == message.source_id || hasPermission(PermissionLevel.Moderator);
    }, [hasPermission, message.source_id, userSteamID]);

    const onDelete = useCallback(async () => {
        try {
            const confirmed = await confirmModal.show({
                title: 'Delete Post?',
                children: (
                    <Box>
                        {isFirstMessage && (
                            <Typography variant={'body1'} fontWeight={700} color={theme.palette.error.dark}>
                                Please be aware that by deleting the first post in the thread, this will result in the deletion of the{' '}
                                <i>entire thread</i>.
                            </Typography>
                        )}
                        <Typography variant={'body1'}>This action cannot be undone.</Typography>
                    </Box>
                )
            });
            if (confirmed) {
                await apiDeleteMessage(activeMessage.forum_message_id);
                onDeleteSuccess(activeMessage.forum_message_id);
            }
            await confirmModal.hide();
            if (isFirstMessage) {
                await navigate({ to: '/forums' });
            }
        } catch (e) {
            logErr(e);
        }
    }, [activeMessage.forum_message_id, confirmModal, isFirstMessage, navigate, onDeleteSuccess, theme.palette.error.dark]);

    return (
        <Paper elevation={1} id={`${activeMessage.forum_message_id}`}>
            <Grid container>
                <Grid xs={2} padding={2} sx={{ backgroundColor: theme.palette.background.paper }}>
                    <Stack alignItems={'center'}>
                        <ForumAvatar
                            alt={activeMessage.personaname}
                            online={activeMessage.online}
                            src={avatarHashToURL(activeMessage.avatarhash, 'medium')}
                        />

                        <ForumRowLink label={activeMessage.personaname} to={`/profile/${activeMessage.source_id}`} align={'center'} />
                        <Typography variant={'subtitle1'} align={'center'}>
                            {permissionLevelString(activeMessage.permission_level)}
                        </Typography>
                    </Stack>
                </Grid>
                <Grid xs={10}>
                    {edit ? (
                        <MessageEditor
                            message={activeMessage}
                            onUpdate={onUpdate}
                            onCancel={() => {
                                setEdit(false);
                            }}
                        />
                    ) : (
                        <Box>
                            <Grid container direction="row" borderBottom={(theme) => theme.palette.divider}>
                                <Grid xs={6}>
                                    <Stack direction={'row'}>
                                        <Typography variant={'body2'} padding={1}>
                                            {renderDateTime(activeMessage.created_on)}
                                        </Typography>
                                        {isAfter(message.created_on, message.updated_on) && (
                                            <Typography variant={'body2'} padding={1}>
                                                {`Edited: ${renderDateTime(activeMessage.updated_on)}`}
                                            </Typography>
                                        )}
                                    </Stack>
                                </Grid>
                                <Grid xs={6}>
                                    <Stack direction="row" justifyContent="end">
                                        <IconButton color={'error'} onClick={onDelete}>
                                            <DeleteForeverIcon />
                                        </IconButton>
                                        {editable && (
                                            <IconButton
                                                title={'Edit Post'}
                                                color={'secondary'}
                                                size={'small'}
                                                onClick={() => {
                                                    setEdit(true);
                                                }}
                                            >
                                                <EditIcon />
                                            </IconButton>
                                        )}
                                        <Typography
                                            padding={1}
                                            component={RouterLink}
                                            variant={'body2'}
                                            to={`#${activeMessage.forum_message_id}`}
                                            textAlign={'right'}
                                            color={(theme: Theme) => {
                                                return theme.palette.text.primary;
                                            }}
                                        >
                                            {`#${activeMessage.forum_message_id}`}
                                        </Typography>
                                    </Stack>
                                </Grid>
                            </Grid>
                            <Grid xs={12} padding={1}>
                                <MarkDownRenderer body_md={activeMessage.body_md} />

                                {activeMessage.signature != '' && (
                                    <>
                                        <Divider />
                                        <MarkDownRenderer body_md={activeMessage.signature} />
                                    </>
                                )}
                            </Grid>
                        </Box>
                    )}
                </Grid>
            </Grid>
        </Paper>
    );
};

// interface MessageEditValues {
//     body_md: string;
// }
//
// const validationSchema = z.object({
//     body_md: bodyMDValidator
// });

const MessageEditor = ({ onCancel }: { message: ForumMessage; onUpdate: (msg: ForumMessage) => void; onCancel: () => void }) => {
    // const onSubmit = useCallback(
    //     async (values: MessageEditValues) => {
    //         try {
    //             const updated = await apiSaveThreadMessage(message.forum_message_id, values.body_md);
    //             onUpdate(updated);
    //         } catch (e) {
    //             logErr(e);
    //         }
    //     },
    //     [message.forum_message_id, onUpdate]
    // );

    return (
        // <Formik<MessageEditValues>
        //     onSubmit={onSubmit}
        //     initialValues={{ body_md: message.body_md }}
        //     validationSchema={validationSchema}
        //     validateOnBlur={true}
        // >
        <Stack padding={1}>
            {/*<MDBodyField />*/}
            <ButtonGroup>
                <Button variant={'contained'} color={'error'} onClick={onCancel}>
                    Cancel
                </Button>
                {/*<SubmitButton />*/}
            </ButtonGroup>
        </Stack>
        // </Formik>
    );
};

function ForumThreadPage() {
    const [updatedMessages, setUpdatedMessages] = useState<ForumMessage[]>();
    const [updatedThread, setUpdatedThread] = useState<ForumThread>();
    const { forum_thread_id } = Route.useParams();
    const { page } = Route.useSearch();

    const { hasPermission, permissionLevel } = useRouteContext({ from: '/_auth/forums/thread/$forum_thread_id' });
    const thread_id = parseInt(forum_thread_id ?? '');
    const navigate = useNavigate();
    const { data: threadOrig } = useThread(thread_id);
    const { data: messagesOrig, count } = useThreadMessages({
        forum_thread_id: thread_id,
        offset: (Number(page) - 1) * RowsPerPage.Ten,
        limit: RowsPerPage.Ten,
        order_by: 'forum_message_id',
        desc: false
    });

    const messages = useMemo(() => {
        return updatedMessages ?? messagesOrig;
    }, [messagesOrig, updatedMessages]);

    const activeThread = useMemo(() => {
        return updatedThread ?? threadOrig;
    }, [threadOrig, updatedThread]);

    useScrollToLocation();

    const firstPostID = useMemo(() => {
        if (Number(page) > 1) {
            return -1;
        }
        if (messages.length > 0) {
            return messages[0].forum_message_id;
        }
        return -1;
    }, [messages, page]);

    const onEditThread = useCallback(async () => {
        try {
            const newThread = await NiceModal.show<ForumThread>(ModalForumThreadEditor, {
                thread: activeThread
            });
            if (newThread.forum_thread_id > 0) {
                setUpdatedThread(newThread);
            } else {
                await navigate({ to: '/forums' });
            }
        } catch (e) {
            logErr(e);
        }
    }, [navigate, activeThread]);

    const onMessageDeleted = useCallback(
        (forum_message_id: number) => {
            setUpdatedMessages(() => {
                return messages.filter((m) => m.forum_message_id != forum_message_id);
            });
        },
        [messages]
    );

    const replyContainer = useMemo(() => {
        if (permissionLevel() == PermissionLevel.Guest) {
            return <LoginPage />;
        } else if (activeThread?.forum_thread_id && !activeThread?.locked) {
            return (
                <ForumThreadReplyBox
                    forum_thread_id={activeThread?.forum_thread_id}
                    onSuccess={(message) => {
                        setUpdatedMessages(() => {
                            return [...messages, message];
                        });
                    }}
                />
            );
        } else {
            return <></>;
        }
    }, [activeThread?.forum_thread_id, activeThread?.locked, messages, permissionLevel]);

    return (
        <Stack spacing={1}>
            <Stack direction={'row'}>
                {hasPermission(PermissionLevel.Moderator) && (
                    <IconButton color={'warning'} onClick={onEditThread}>
                        <ConstructionIcon fontSize={'small'} />
                    </IconButton>
                )}
                <Typography variant={'h3'}>{activeThread?.title}</Typography>
            </Stack>
            <Stack direction={'row'} spacing={1}>
                <Person2Icon />
                <VCenterBox>
                    <Typography
                        variant={'body2'}
                        component={RouterLink}
                        color={(theme: Theme) => theme.palette.text.primary}
                        to={`/profile/${activeThread?.source_id}`}
                    >
                        {activeThread?.personaname ?? ''}
                    </Typography>
                </VCenterBox>
                <AccessTimeIcon />
                <VCenterBox>
                    <Typography variant={'body2'}>{renderDateTime(activeThread?.created_on ?? new Date())}</Typography>
                </VCenterBox>
            </Stack>
            {messages.map((m) => (
                <ThreadMessageContainer
                    message={m}
                    key={`thread-message-id-${m.forum_message_id}`}
                    onDeleteSuccess={onMessageDeleted}
                    isFirstMessage={firstPostID == m.forum_message_id}
                />
            ))}
            <Pagination
                count={count > 0 ? Math.ceil(count / RowsPerPage.Ten) : 0}
                page={Number(page)}
                onChange={async (_, newPage) => {
                    await navigate({ search: (prev) => ({ ...prev, page: newPage }) });
                }}
            />
            {activeThread?.locked && (
                <Paper>
                    <Typography variant={'h4'} textAlign={'center'} padding={1}>
                        <LockIcon /> Thread Locked
                    </Typography>
                </Paper>
            )}
            {replyContainer}
        </Stack>
    );
}
