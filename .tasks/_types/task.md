---
kind: mdbase.type
name: task
version: 1
description: A task managed by TaskNotes.
match:
  where:
    tags:
      contains: task
schema:
  dialect: json-schema-2020-12
  value:
    $schema: https://json-schema.org/draft/2020-12/schema
    type: object
    additionalProperties: true
    properties:
      id:
        type: string
        minLength: 1
      title:
        type: string
        minLength: 1
      status:
        enum: &a1
          - open
          - in-progress
          - done
          - cancelled
        default: open
      priority:
        enum: &a3
          - low
          - normal
          - high
          - urgent
        default: normal
      due:
        anyOf:
          - type: string
            format: date
          - type: string
            format: date-time
      scheduled:
        anyOf:
          - type: string
            format: date
          - type: string
            format: date-time
      contexts:
        type: array
        items:
          type: string
      projects:
        type: array
        items:
          type: string
      timeEstimate:
        type: integer
        minimum: 0
      completedDate:
        type: string
        format: date
      dateCreated:
        type: string
        format: date-time
      dateModified:
        type: string
        format: date-time
      recurrence:
        type: string
      recurrenceAnchor:
        enum:
          - scheduled
          - completion
        default: scheduled
      occurrenceMaterialization:
        enum:
          - manual
          - on_completion
          - rolling
        default: manual
      occurrenceNextTrigger:
        enum:
          - completion
          - completion_or_skip
        default: completion
      occurrenceTemplate:
        type: string
      occurrencePastHorizon:
        type: string
      occurrenceFutureHorizon:
        type: string
      recurrenceParent:
        type: string
      occurrenceDate:
        type: string
        format: date
      tags:
        type: array
        items:
          type: string
      timeEntries:
        type: array
        items:
          type: object
          additionalProperties: false
          properties:
            startTime:
              type: string
              format: date-time
            endTime:
              type: string
              format: date-time
            description:
              type: string
            duration:
              type: integer
      reminders:
        type: array
        items:
          oneOf:
            - type: object
              required:
                - id
                - type
                - absoluteTime
              additionalProperties: false
              properties:
                id:
                  type: string
                type:
                  const: absolute
                description:
                  type: string
                absoluteTime:
                  type: string
                  format: date-time
            - type: object
              required:
                - id
                - type
                - relatedTo
                - offset
              additionalProperties: false
              properties:
                id:
                  type: string
                type:
                  const: relative
                description:
                  type: string
                relatedTo:
                  enum:
                    - due
                    - scheduled
                offset:
                  type: string
      blockedBy:
        type: array
        items:
          type: object
          additionalProperties: false
          properties:
            uid:
              type: string
            reltype:
              type: string
            gap:
              type: string
          required:
            - uid
      completeInstances:
        type: array
        items:
          type: string
          format: date
      skippedInstances:
        type: array
        items:
          type: string
          format: date
      icsEventId:
        type: array
        items:
          type: string
      googleCalendarEventId:
        type: string
      googleCalendarExceptionEventId:
        type: string
      googleCalendarExceptionOriginalScheduled:
        type: string
        format: date
      googleCalendarMovedOriginalDates:
        type: array
        items:
          type: string
          format: date
      tasknotes_manual_order:
        type: number
    allOf:
      - if:
          required:
            - status
          properties:
            status:
              enum: &a2
                - done
          not:
            required:
              - recurrence
        then:
          required:
            - completedDate
    required:
      - title
      - status
      - dateCreated
collection:
  read_defaults:
    status: open
    priority: normal
    recurrenceAnchor: scheduled
    occurrenceMaterialization: manual
    occurrenceNextTrigger: completion
  links:
    projects[]:
      target_type: any
      validate_exists: false
    occurrenceTemplate:
      target_type: any
      validate_exists: false
    recurrenceParent:
      target_type: task
      validate_exists: false
    blockedBy[].uid:
      target_type: task
      validate_exists: false
  path:
    runtime: tasknotes
    template: "{{title}}"
    folder: tasks
    generated_by: tasknotes.filename.create
  display:
    name_field: title
  unique:
    - field: id
      scope: type
lifecycle:
  on_create:
    set:
      id:
        uuid: true
      dateCreated:
        now: true
      dateModified:
        now: true
  on_update:
    set:
      dateModified:
        now: true
implements:
  - contract: tasknotes.task
    version: 0.3.0-rc.1
    fields:
      id: id
      title: title
      status: status
      priority: priority
      due: due
      scheduled: scheduled
      contexts: contexts
      projects: projects
      timeEstimate: timeEstimate
      completedDate: completedDate
      dateCreated: dateCreated
      dateModified: dateModified
      recurrence: recurrence
      recurrenceAnchor: recurrenceAnchor
      occurrenceMaterialization: occurrenceMaterialization
      occurrenceNextTrigger: occurrenceNextTrigger
      occurrenceTemplate: occurrenceTemplate
      occurrencePastHorizon: occurrencePastHorizon
      occurrenceFutureHorizon: occurrenceFutureHorizon
      recurrenceParent: recurrenceParent
      occurrenceDate: occurrenceDate
      tags: tags
      timeEntries: timeEntries
      reminders: reminders
      blockedBy: blockedBy
      completeInstances: completeInstances
      skippedInstances: skippedInstances
      icsEventId: icsEventId
      googleCalendarEventId: googleCalendarEventId
      googleCalendarExceptionEventId: googleCalendarExceptionEventId
      googleCalendarExceptionOriginalScheduled: googleCalendarExceptionOriginalScheduled
      googleCalendarMovedOriginalDates: googleCalendarMovedOriginalDates
      sortOrder: tasknotes_manual_order
    binding:
      profiles:
        - core-lite
        - recurrence
        - templating
        - materialized-occurrences
        - extended
      capabilities:
        - dependencies
        - reminders
        - links
        - time-tracking
        - materialized-occurrences
        - archive
        - templating
      title:
        storage: frontmatter
        filename_format: title
      status:
        values: *a1
        default: open
        completed_values: *a2
        skipped_values:
          - cancelled
        default_skipped: cancelled
        definitions:
          - value: open
            label: open
            color: ""
            is_completed: false
            is_skipped: false
            exclude_from_cycle: false
            order: 0
            auto_archive: false
            auto_archive_delay_minutes: 0
          - value: in-progress
            label: in-progress
            color: ""
            is_completed: false
            is_skipped: false
            exclude_from_cycle: false
            order: 1
            auto_archive: false
            auto_archive_delay_minutes: 0
          - value: done
            label: done
            color: ""
            is_completed: true
            is_skipped: false
            exclude_from_cycle: false
            order: 2
            auto_archive: false
            auto_archive_delay_minutes: 0
          - value: cancelled
            label: cancelled
            color: ""
            is_completed: false
            is_skipped: true
            exclude_from_cycle: false
            order: 3
            auto_archive: false
            auto_archive_delay_minutes: 0
      priority:
        values: *a3
        default: normal
        definitions:
          - value: low
            label: low
            color: ""
            weight: 0
          - value: normal
            label: normal
            color: ""
            weight: 1
          - value: high
            label: high
            color: ""
            weight: 2
          - value: urgent
            label: urgent
            color: ""
            weight: 3
      recurrence:
        syntax: tasknotes
        maintain_due_date_offset: true
        reset_body_checkboxes: false
      occurrences:
        identity_roles:
          - recurrenceParent
          - occurrenceDate
        default_materialization: manual
        default_next_trigger: completion
        past_horizon: P0D
        future_horizon: P14D
      links:
        accepted_formats:
          - wikilink
          - markdown
        write_format: wikilink
      archive:
        archived_tag: archived
        move_on_archive: false
      time_tracking:
        auto_stop_on_complete: false
      templating:
        enabled: false
        occurrence_enabled: false
x-tasknotes-generator:
  managed_fields:
    - blockedBy
    - completeInstances
    - completedDate
    - contexts
    - dateCreated
    - dateModified
    - due
    - googleCalendarEventId
    - googleCalendarExceptionEventId
    - googleCalendarExceptionOriginalScheduled
    - googleCalendarMovedOriginalDates
    - icsEventId
    - id
    - occurrenceDate
    - occurrenceFutureHorizon
    - occurrenceMaterialization
    - occurrenceNextTrigger
    - occurrencePastHorizon
    - occurrenceTemplate
    - priority
    - projects
    - recurrence
    - recurrenceAnchor
    - recurrenceParent
    - reminders
    - scheduled
    - skippedInstances
    - status
    - tags
    - tasknotes_manual_order
    - timeEntries
    - timeEstimate
    - title
---

# Task

This type definition implements the TaskNotes contract for this mdbase collection.
Its JSON Schema describes persisted task frontmatter; collection and lifecycle
metadata describe generic mdbase behavior; `implements` maps the portable
TaskNotes task view and supplies TaskNotes behavior.

Changes made here are loaded by TaskNotes. Portable changes made in TaskNotes
settings are written back while unknown extensions are preserved.
