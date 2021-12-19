import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:control_center/page.dart';
import 'package:dbus/dbus.dart';

class LenovoSettingsPage extends SettingsPage {
  const LenovoSettingsPage({Key? key, required DBusClient client})
      : super(key: key, title: 'Lenovo', client: client);

  @override
  Widget buildPage(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: SettingCategory(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.start,
          children: <Widget>[
            Text("Conservation Mode",
                style: Theme.of(context).textTheme.headline5),
            const Padding(padding: EdgeInsets.symmetric(vertical: 8.0)),
            LenovoConservationModeSetting(
                client: client,
                showDBusError: (text) => showDBusError(context, text)),
          ],
        ),
      ),
    );
  }
}

class LenovoConservationModeSetting extends StatefulWidget {
  const LenovoConservationModeSetting(
      {Key? key, required this.client, required this.showDBusError})
      : super(key: key);

  final DBusClient client;
  final Future<void> Function(List<Widget>) showDBusError;

  @override
  createState() => _LenovoConservationModeSettingState();
}

class _LenovoConservationModeSettingState
    extends State<LenovoConservationModeSetting> {
  bool enabled = false;
  bool available = false;
  late Timer _timer;

  Future<void> updateStateFromDBus() async {
    var object = DBusRemoteObject(widget.client,
        name: 'com.retrixe.ControlCenter.v0',
        path: DBusObjectPath('/com/retrixe/ControlCenter/v0'));
    var result = await object.callMethod(
        'com.retrixe.ControlCenter.v0', 'LenovoGetConservationModeStatus', []);
    num value = result.values[0].toNative();
    setState(() {
      enabled = value == 1;
      available = value == 0 || value == 1;
    });
  }

  Future<void> setConservationMode(bool value) async {
    var object = DBusRemoteObject(widget.client,
        name: 'com.retrixe.ControlCenter.v0',
        path: DBusObjectPath('/com/retrixe/ControlCenter/v0'));
    await object.callMethod('com.retrixe.ControlCenter.v0',
        'LenovoSetConservationMode', [DBusBoolean(value)]);
  }

  @override
  void initState() {
    super.initState();
    updateStateFromDBus().catchError((error) {
      stderr.writeln(error);
      widget.showDBusError(const [
        Text("An error occurred when talking to the Control Center daemon."),
        Text("The app WILL not work correctly!"),
      ]);
    });
    _timer = Timer.periodic(const Duration(seconds: 5), (timer) async {
      try {
        await updateStateFromDBus();
      } catch (error) {
        stderr.writeln(error);
        widget.showDBusError(const [
          Text("An error occurred when talking to the Control Center daemon."),
          Text("The app WILL not work correctly!"),
        ]);
      }
    });
  }

  @override
  void dispose() {
    super.dispose();
    _timer.cancel();
  }

  void toggleHandler(bool value) {
    setState(() => enabled = value);
    setConservationMode(value).catchError((error) {
      setState(() => enabled = !value);
      stderr.writeln(error);
      widget.showDBusError(const [
        Text(
            "Control Center encountered an error when trying to set this setting."),
        Text(
            "The setting might not be applied correctly. Check if the control center daemon is running."),
      ]);
    });
  }

  @override
  Widget build(BuildContext context) {
    if (!available) {
      return Text('Not available on this device.',
          style: Theme.of(context).textTheme.subtitle1);
    }
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text('Enabled', style: Theme.of(context).textTheme.subtitle1),
        Switch(
          value: enabled,
          onChanged: available ? toggleHandler : null,
        )
      ],
    );
  }
}
